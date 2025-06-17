# photon-cli

**Photon CLI** is the official command-line interface for the [Photon](https://github.com/smtdfc/photon) backend framework. It streamlines the development process by scaffolding applications, modules, and other project components with simple commands.

---

## 🚀 Features

* Generate new Photon-based applications
* Create and register modules with boilerplate setup
* Consistent project structure
* Developer-friendly and easy to extend

---

## 📦 Installation

```bash
go install github.com/smtdfc/photon-cli@latest
```

Make sure `$GOPATH/bin` is in your system's `PATH` to run `photon` from anywhere.

---

## 🛠️ Commands

### `photon gen app <name>`

Scaffold a new Photon project in a folder named `<name>`.

```bash
photon gen app my-app
```

### `photon gen module <name>`

Generate a module with routing and handler boilerplate.

```bash
photon gen module user
```

---

## 📄 Example Workflow

```bash
# Step 1: Create a new app
go mod init example.com/hello
photon gen app hello

# Step 2: Generate a module
photon gen module hello

# Step 3: Run your Photon server
go run .
```

Visit: [http://localhost:3000/hello](http://localhost:3000/hello)

---

## 📚 Extending CLI

Photon CLI is designed to be extensible. You can add custom generators by modifying the codebase or contributing to the project.

---

## 🧑‍💻 Contributing

Contributions are welcome! Open an issue or submit a PR to help improve `photon-cli`.

GitHub: [https://github.com/smtdfc/photon-cli](https://github.com/smtdfc/photon-cli)

---

## 📜 License

MIT © 2025 [smtdfc](https://github.com/smtdfc)
