# Práctica Guiada 2 – Sistemas Distribuidos

## 🧠 Descripción
Este proyecto implementa una aplicación cliente-servidor en Go usando **gRPC**, en la que múltiples clientes envían *heartbeats* cada 5 segundos, y el servidor detecta si alguno de ellos falla o deja de enviar señales.

---

## 👥 Integrantes del grupo
- Amarilla, Aldo
- Stigelmeier, Marcela

---

## 🛠️ Instrucciones para compilar y ejecutar

### 1. Generar el código desde el archivo `.proto`

Desde la raíz del proyecto:

```bash
protoc --go_out=. --go-grpc_out=. proto/monitor.proto
