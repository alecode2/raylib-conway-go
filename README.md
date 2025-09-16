# Conway's Game of Life (Go + Raylib)  

This project is an implementation of the classic **Conway's Game of Life**, built in **Golang** using the [Raylib](https://www.raylib.com/) graphics library.  

The most interesting aspect of the project is not just the simulation itself, but the creation of a **custom UI library**, featuring a **flexbox-based layout system**. This allows for building user interfaces in a declarative and flexible way, inspired by modern frontend frameworks.  

## ✨ Features

- Full implementation of **Conway’s Game of Life** with real-time rendering using **Raylib**.  
- **Custom UI library in Go**:  
  - Layout system based on **flexbox**, similar to CSS.  
  - Support for nested components and dynamic layouts.  
  - Basic UI elements (panels, buttons, etc.) built on top of the layout engine.  
- Modular and extensible codebase — easy to expand with new components or simulations.  

## 📂 Project structure 
```
├── assets/ # Asset management (loading/unloading resources)
├── event/ # Global event bus system
├── game/ # Conway's Game of Life core logic
├── input/ # Input handling + event firing for input events
├── render/ # Rendering logic
├── ui/ # Custom UI library
│ └── components/ # UI components built on top of the flexbox system
├── README.md
├── env.sh # Environment setup script
├── go.mod # Go modules definition
├── go.sum # Go modules checksums
├── main.go # Entry point
└── run.sh # Helper script to run the project
```
🎯 Purpose

This project goes beyond implementing the Game of Life — it demonstrates how to design reusable tools from scratch in Go:
- Graphics programming with Raylib.
- Component-based UI architecture.
- Flexbox-inspired layout system, bridging concepts from web development into a native application.
- Event-driven design, with a global event bus for decoupled communication.

📸 Screenshots

WIP
