<div align="center">

# ✦ Mawey Bacelonia — SSH Portfolio

**A terminal-based portfolio built with Go, Bubble Tea, and Lip Gloss.**

Instead of a traditional web portfolio, this project explores building an
interactive portfolio experience directly in the terminal.



<img src="assets/demo.gif" alt="SSH Portfolio Demo" width="700">


[![Go](https://img.shields.io/badge/Go-1.26-00ADD8?style=flat-square&logo=go&logoColor=white)](https://go.dev)
[![Bubble Tea](https://img.shields.io/badge/Bubble%20Tea-TUI-FF69B4?style=flat-square)](https://github.com/charmbracelet/bubbletea)
[![Lip Gloss](https://img.shields.io/badge/Lip%20Gloss-Styling-9B59B6?style=flat-square)](https://github.com/charmbracelet/lipgloss)
[![Docker](https://img.shields.io/badge/Docker-Ready-2496ED?style=flat-square&logo=docker&logoColor=white)](https://www.docker.com)
[![License](https://img.shields.io/badge/License-MIT-lightgrey?style=flat-square)](#)

</div>

## ✦ Features

| | |
|---|---|
| 🧭 | Interactive terminal navigation |
| 📐 | Responsive layouts for different terminal sizes |
| 🖼️ | ASCII portrait with dark / light variants |
| ⌨️ | Full keyboard navigation |
| ✉️ | Contact page |
| 🐳 | Dockerized application |

<br>

## 🛠 Built With

<div align="left">

![Go](https://img.shields.io/badge/-Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![Bubble Tea](https://img.shields.io/badge/-Bubble%20Tea-FF69B4?style=for-the-badge)
![Lip Gloss](https://img.shields.io/badge/-Lip%20Gloss-9B59B6?style=for-the-badge)
![Docker](https://img.shields.io/badge/-Docker-2496ED?style=for-the-badge&logo=docker&logoColor=white)

</div>

<br>

## 🚀 Running Locally

**Clone the repository**

```bash
git clone https://github.com/MaweySB02/ssh-portfolio.git
cd ssh-portfolio
```

**Run with Go**

```bash
go run .
```

**Run with Docker Compose**

```bash
docker compose run --rm portfolio
```

<br>

## 🎨 Controls

| Key | Action |
|:---:|---|
| `↑` / `↓` | Navigate |
| `j` / `k` | Navigate |
| `Enter` | Select |
| `Esc` / `Backspace` | Return to menu |
| `t` | Toggle ASCII portrait |
| `q` | Quit |

<br>

## 📐 Responsive Layout

The portfolio adapts its layout based on terminal width:

| Width | Layout |
|---|---|
| `120+` columns | Full layout |
| `80–119` columns | Stacked layout |
| `< 80` columns | Compact layout |

<br>

## 📁 Project Structure

```
ssh-portfolio/
├── main.go
├── Dockerfile
├── docker-compose.yml
├── .dockerignore
├── go.mod
├── go.sum
└── .gitignore
```

<br>

## 👤 About

<div align="center">

Built by **Mawey Bacelonia**, a Software Developer interested in creating
software that feels creative and human.

</div>