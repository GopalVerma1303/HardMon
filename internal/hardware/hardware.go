package hardware

import (
	"fmt"
	"runtime"
	"strings"
	"time"

	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/disk"
	"github.com/shirou/gopsutil/v4/host"
	"github.com/shirou/gopsutil/v4/load"
	"github.com/shirou/gopsutil/v4/mem"
	"github.com/shirou/gopsutil/v4/net"
)

const (
	megabyteDiv uint64 = 1024 * 1024
	gigabyteDiv uint64 = megabyteDiv * 1024
)

func GetSystemSection() (string, error) {
	runTimeOS := runtime.GOOS
	vmStat, err := mem.VirtualMemory()
	if err != nil {
		return "", err
	}
	hostStat, err := host.Info()
	if err != nil {
		return "", err
	}
	uptime, _ := host.Uptime()
	bootTime, _ := host.BootTime()
	loadAvg, _ := load.Avg()

	html := "<div class='system-data'><table class='table table-striped table-hover table-sm'><tbody>"
	html += fmt.Sprintf("<tr><td>Operating System:</td> <td><i class='fa fa-brands fa-%s'></i> %s</td></tr>", getOSIcon(runTimeOS), runTimeOS)
	html += fmt.Sprintf("<tr><td>Platform:</td><td> <i class='fa fa-brands fa-%s'></i> %s</td></tr>", getPlatformIcon(hostStat.Platform), hostStat.Platform)
	html += fmt.Sprintf("<tr><td>Hostname:</td><td>%s</td></tr>", hostStat.Hostname)
	html += fmt.Sprintf("<tr><td>Kernel Version:</td><td>%s</td></tr>", hostStat.KernelVersion)
	html += fmt.Sprintf("<tr><td>Uptime:</td><td>%s</td></tr>", formatUptime(uptime))
	html += fmt.Sprintf("<tr><td>Boot Time:</td><td>%s</td></tr>", time.Unix(int64(bootTime), 0).Format("2006-01-02 15:04:05"))
	html += fmt.Sprintf("<tr><td>Number of processes running:</td><td>%d</td></tr>", hostStat.Procs)
	html += fmt.Sprintf("<tr><td>Total memory:</td><td>%d MB</td></tr>", vmStat.Total/megabyteDiv)
	html += fmt.Sprintf("<tr><td>Used memory:</td><td>%d MB</td></tr>", vmStat.Used/megabyteDiv)
	html += fmt.Sprintf("<tr><td>Free memory:</td><td>%d MB</td></tr>", vmStat.Free/megabyteDiv)
	html += fmt.Sprintf("<tr><td>Percentage used memory:</td><td>%.2f%%</td></tr>", vmStat.UsedPercent)
	html += fmt.Sprintf("<tr><td>Load Average (1, 5, 15 min):</td><td>%.2f, %.2f, %.2f</td></tr>", loadAvg.Load1, loadAvg.Load5, loadAvg.Load15)
	html += "</tbody></table></div>"
	return html, nil
}

func GetDiskSection() (string, error) {
	partitions, err := disk.Partitions(false)
	if err != nil {
		return "", err
	}

	html := "<div class='disk-data'><table class='table table-striped table-hover table-sm'><tbody>"
	html += "<tr><th>Mount Point</th><th>Total</th><th>Used</th><th>Free</th><th>Usage</th><th>File System</th></tr>"

	for _, partition := range partitions {
		usage, err := disk.Usage(partition.Mountpoint)
		if err != nil {
			continue
		}
		html += fmt.Sprintf("<tr><td>%s</td><td>%d GB</td><td>%d GB</td><td>%d GB</td><td>%.2f%%</td><td>%s</td></tr>",
			partition.Mountpoint,
			usage.Total/gigabyteDiv,
			usage.Used/gigabyteDiv,
			usage.Free/gigabyteDiv,
			usage.UsedPercent,
			partition.Fstype)
	}

	html += "</tbody></table></div>"
	return html, nil
}

func GetCpuSection() (string, error) {
	cpuStat, err := cpu.Info()
	if err != nil {
		return "", err
	}
	percentage, err := cpu.Percent(time.Second, true)
	if err != nil {
		return "", err
	}

	html := "<div class='cpu-data'><table class='table table-striped table-hover table-sm'><tbody>"
	if len(cpuStat) != 0 {
		html += fmt.Sprintf("<tr><td>Model Name:</td><td>%s</td></tr>", cpuStat[0].ModelName)
		html += fmt.Sprintf("<tr><td>Vendor:</td><td>%s</td></tr>", cpuStat[0].VendorID)
		html += fmt.Sprintf("<tr><td>Family:</td><td>%s</td></tr>", cpuStat[0].Family)
		html += fmt.Sprintf("<tr><td>Number of Cores:</td><td>%d</td></tr>", cpuStat[0].Cores)
		html += fmt.Sprintf("<tr><td>Clock Speed:</td><td>%.2f MHz</td></tr>", cpuStat[0].Mhz)
		html += fmt.Sprintf("<tr><td>Cache Size:</td><td>%d KB</td></tr>", cpuStat[0].CacheSize)
	}

	html += "<tr><td>CPU Usage:</td><td><div class='row'>"
	for idx, cpupercent := range percentage {
		html += fmt.Sprintf("<div class='col-md-3 mb-2'>CPU %d: %.2f%%</div>", idx, cpupercent)
	}
	html += "</div></td></tr></tbody></table></div>"

	return html, nil
}

func GetNetworkSection() (string, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}

	html := "<div class='network-data'><table class='table table-striped table-hover table-sm'><tbody>"
	html += "<tr><th>Interface</th><th>IP Addresses</th><th>MAC Address</th><th>Flags</th></tr>"

	for _, iface := range interfaces {
		var ips []string
		for _, addr := range iface.Addrs {
			ips = append(ips, addr.Addr)
		}

		html += fmt.Sprintf("<tr><td>%s</td><td>%s</td><td>%s</td><td>%s</td></tr>",
			iface.Name,
			strings.Join(ips, ", "),
			iface.HardwareAddr,
			strings.Join(iface.Flags, ", "))
	}

	html += "</tbody></table></div>"
	return html, nil
}

func formatUptime(uptime uint64) string {
	days := uptime / (60 * 60 * 24)
	hours := (uptime % (60 * 60 * 24)) / (60 * 60)
	minutes := (uptime % (60 * 60)) / 60
	return fmt.Sprintf("%d days, %d hours, %d minutes", days, hours, minutes)
}

func getOSIcon(os string) string {
	switch os {
	case "linux":
		return "linux"
	case "windows":
		return "windows"
	case "darwin":
		return "apple"
	default:
		return "desktop"
	}
}

func getPlatformIcon(platform string) string {
	switch strings.ToLower(platform) {
	case "ubuntu":
		return "ubuntu"
	case "fedora":
		return "fedora"
	case "centos":
		return "centos"
	case "debian":
		return "debian"
	default:
		return "linux"
	}
}
