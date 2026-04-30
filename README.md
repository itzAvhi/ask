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

Run the commnad given below for the most easiest download.

```bash
    # 1. Download the binary directly from your release
wget https://github.com/itzAvhi/ask/releases/download/v1.0.0/ask-linux-amd64

# 2. Make the file executable
chmod +x ask-linux-amd64

# 3. Run the setup and first query
./ask-linux-amd64 "how do I use this tool?"
```

---
# Configuration
The first time you run the tool, it will automatically startup the setup

1. Setup

The terminal will ask: `Enter Your api key: `. You will need a api key to make `ask` work. Get the free api at `https://console.groq.com/`. (*NOTE: YOU CAN ONLY PLUG IN THE GROQ API KEY, OTHERS WONT WORK*)

---
You can now sucessfully use `ask`

# Usage Examples 
- `ask "what is my local ip address"`
- `ask "undo my last commit but keep changes"`

*ProTip: ypu can use ask without typing `""` (Double inverted comma)*
