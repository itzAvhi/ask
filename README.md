# Ask - AI Powered Terminal Helper 

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![Arch Linux](https://img.shields.io/badge/Arch%20Linux-1793D1?style=for-the-badge&logo=arch-linux&logoColor=white)
![Llama3](https://img.shields.io/badge/Llama_3-0467DF?style=for-the-badge&logo=meta&logoColor=white)

`ask` is a minimalistic, insanely fast CLI tool that uses AI(Groq/Llama 3) to convert natural language querries into executiable terminal commands. It features an automated wisard triggred upon the first querry and is designed specifically for Linux


## Why `ask`?

Standard AI assistants don't know your environment. `ask` bridges the gap by injecting real-time system context into every query.

| Feature | Standard AI | `ask` (Your Tool) |
| :--- | :--- | :--- |
| **Path Knowledge** | Guesses (e.g., `/home/user`) | **Real** (e.g., `/home/avhi/ask`) |
| **File Awareness** | None | **Real-time** (`ls -F` injected) |
| **Package Manager** | Generic (`apt`/`brew`) | **Specific** (Detects `pacman`) |
| **Navigation** | Fails (Subshell limitation) | **Works** (`os.Chdir` implementation) |
| **OS Identity** | Generic Linux | **Deep** (Detects Arch Linux/Hyprland) |

---

![Ask Cli demo](/home/avhi/askdemo1.png)
![Ask Cli demo](/home/avhi/askdemo2.png)
![Ask Cli demo](/home/avhi/askdemo3.png)



## Full Linux Setup

This guide covers everything you need to know to install and use this tool globally.

### 1. Install Go (The compiler )
if you dont have Go installed, Run the following command based on your distribution

**Arch Linux**
```bash
    sudo pacman -S go
```

**Ubuntu/Debian**
```bash
    sudo apt update && sudo apt install golang
```

**Verify the installation**
```bash
    go version
```

### 2.Clone and Build
Clone the repository and compile the binary:
```bash
    git clone https://github.com/itzAvhi/ask.git && cd ask && go build -o ask main.go
```
### Make it Global
to run `ask` from any folder without typing the full path, move it to your system's binary folder
```bash
    sudo mv ask /usr/local/bin/
```
You will need a api key to make `ask` work. Get the free api at `https://console.groq.com/`. (*NOTE: YOU CAN ONLY PLUG IN THE GROQ API KEY, OTHERS WONT WORK*)

### Then run this if you use bash
```bash
echo "export GROQ_API_KEY='your_actual_key_here'" >> ~/.bashrc
```
### OR this for zsh users
```
echo "export GROQ_API_KEY='your_key'" >> ~/.zshrc && source ~/.zshrc
```
*Note: you see 'your api key' placeholder. replace it with your acctual api key*
---
# Configuration
The first time you run the tool, it will automatically startup the setup

### 1. Run a querry
```bash
   ask "command to know what is my name"
```
### 2. Setup

The terminal will ask: `Enter Your api key: `. You will need a api key to make `ask` work. Get the free api at `https://console.groq.com/`. (*NOTE: YOU CAN ONLY PLUG IN THE GROQ API KEY, OTHERS WONT WORK*)

---
You can now sucessfully use `ask`

# Usage Examples 
- `ask "what is my local ip address"`
- `ask "undo my last commit but keep changes"`

[![Ask - A terminal Buddy](https://img.youtube.com/vi/16DW1NL42JU/maxresdefault.jpg)](https://www.youtube.com/watch?v=16DW1NL42JU)

[Watch the full demo of Ask - A terminal Buddy here!](https://www.youtube.com/watch?v=16DW1NL42JU)

*ProTip: ypu can use ask without typing `""` (Double inverted comma)*
*For reviewers: Sorry for the multiple update readme.md command (i had never written a readme in myself this was the first time ;))*
