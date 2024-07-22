<div align="center">

<div align="center">
  <img src="https://github.com/GopalVerma1303/HardMon/blob/872c5967d7568b56ce4288eb4bc74f87432425f7/hm.png" alt="AutoForge Logo" height="100">
</div>

# HardMon

Hardmon is a powerful hardware monitoring tool built with Go and HTMX, providing real-time system metrics with a responsive web interface.

<br />

[![GitHub stars](https://img.shields.io/github/stars/GopalVerma1303/HardMon.svg?style=social&label=Star)](https://github.com/GopalVerma1303/HardMon)
[![GitHub forks](https://img.shields.io/github/forks/GopalVerma1303/HardMon.svg?style=social&label=Fork)](https://github.com/GopalVerma1303/HardMon/fork)
[![GitHub watchers](https://img.shields.io/github/watchers/GopalVerma1303/HardMon.svg?style=social&label=Watch)](https://github.com/GopalVerma1303/HardMon)

</div>

## 🚀 Features

- Real-time hardware monitoring
- Lightweight and efficient Go backend
- Dynamic UI updates with HTMX
- Cross-platform compatibility
- Customizable metrics display
- Dark/Light mode toggle

## 🛠️ Tech Stack

- [Go](https://golang.org/)
- [HTMX](https://htmx.org/)
- [WebSockets](https://developer.mozilla.org/en-US/docs/Web/API/WebSockets_API)
- [Bootstrap](https://getbootstrap.com/)

## 🚀 Getting Started

1. Clone the repository:
```
git clone https://github.com/YourUsername/HardMon.git
```

2. Navigate to the project directory:
```
cd HardMon
```

3. Run the Go server:
```
go run cmd/main.go
```

4. Open `http://localhost:8080` in your browser (or the port specified in your configuration).

## ⚙️ How it works

The tool is typically integrated into a system by running it as a binary executable. Here's a quick overview of the process:

 1. Compile the Go code into a binary executable.
 2. Place the binary and any necessary static files (like HTML templates) on the target system.
 3. Run the binary, which starts a web server (usually on port 8080).
 4. The tool then collects system information in real-time using the gopsutil library.
 5. Access the monitoring dashboard by opening a web browser and navigating to http://localhost:8080 (or the appropriate IP/port).

The binary can be run manually, set up as a system service (e.g., systemd on Linux), or configured to start automatically on boot. This allows for flexible deployment and integration into various system management workflows.

## Screenshots

| Light | Dark |
|----------|----------|
| <img src="https://github.com/GopalVerma1303/HardMon/blob/93eec0422fa2332e6a1821df49371c68ab2fefb0/ss-light.png" /> | <img src="https://github.com/GopalVerma1303/HardMon/blob/93eec0422fa2332e6a1821df49371c68ab2fefb0/ss-dark.png" /> |

| Section | SS |
|----------|----------|
| System | <img src="https://github.com/GopalVerma1303/HardMon/blob/93eec0422fa2332e6a1821df49371c68ab2fefb0/system.png" /> |
| CPU | <img src="https://github.com/GopalVerma1303/HardMon/blob/93eec0422fa2332e6a1821df49371c68ab2fefb0/cpu.png" /> |
| Disk | <img src="https://github.com/GopalVerma1303/HardMon/blob/93eec0422fa2332e6a1821df49371c68ab2fefb0/disk.png" /> |
| Network | <img src="https://github.com/GopalVerma1303/HardMon/blob/93eec0422fa2332e6a1821df49371c68ab2fefb0/network.png" /> | 

## 📝 Documentation

For detailed documentation on each component and usage instructions, please refer to the `docs` folder.

## 🤝 Contributing

Contributions are welcome! Please check out our [Contributing Guide](CONTRIBUTING.md) for more details.

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgements

- [Go](https://golang.org/) for the backend runtime
- [HTMX](https://htmx.org/) for dynamic HTML
- [Bootstrap](https://getbootstrap.com/) for responsive design
- [Font Awesome](https://fontawesome.com/) for icons

---

Made with ❤️ by [Gopal Verma](https://github.com/GopalVerma1303)
