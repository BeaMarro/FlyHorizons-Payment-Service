# 💳 FlyHorizons - Payment Service

This is the **Payment Service** for **FlyHorizons**, an enterprise-grade airline booking system. The service is responsible for handling mock payments, processing transactions, and validating payment events related to flight bookings.

---

## 🚀 Overview

This microservice manages the **payment lifecycle** within the FlyHorizons system. It integrates with the Booking Service via **RabbitMQ**, processes transactions, and validates payment requests. Built with **Go** and the **Gin** framework, it connects to an **Azure-hosted Microsoft SQL Server** database for transaction persistence.

---

## 🛠️ Tech Stack

- **Language**: Go (Golang)
- **Framework**: Gin
- **Database**: Microsoft SQL Server (Azure Hosted)
- **Messaging**: RabbitMQ
- **Architecture**: Microservices

---

## 📦 Features

- 💳 **Mock Payment Handling** for development and testing environments
- 🧾 **Transaction Validation** to simulate payment outcomes
- 🔄 **Processes Events** from the Booking Service
- 📬 **Publishes Events** on payment status (success, failure)
- 🧩 **Service Integration** with Booking via RabbitMQ
- ⚠️ **Centralized Error Handling**

---

## 📄 License
This project is shared for educational and portfolio purposes only. Commercial use, redistribution, or modification is not allowed without explicit written permission. All rights reserved © 2025 Beatrice Marro.

## 👤 Author
Beatrice Marro GitHub: https://github.com/beamarro
