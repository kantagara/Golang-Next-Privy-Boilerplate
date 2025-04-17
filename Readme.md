# Golang + MongoDB + NextJS + Privy Auth Boilerplate

🚀 A clean and minimal starter template for building fullstack apps with:

- ✅ [Golang](https://golang.org/) and [Gin](https://github.com/gin-gonic/gin) for routing
- ✅ [MongoDB](https://www.mongodb.com/) for user storage
- ✅ [Privy](https://www.privy.io/) for web3 & email-based authentication
- ✅ JWT verification using Privy-issued identity tokens (ES256)
- ✅ Clean layered architecture (Handler → Service → Repository)
- ✅ Next.js + TypeScript frontend

---


## Motivation
The main motivation behind this was to create a boilerplate that is ready to use, that can verify the correctness of the identity token on the backend, and extract necessary information (in the case of this repo, that's the wallet address) from the identity token since that part was missing in the Privy documentation for any non-JS/TS language.

---

## 🔧 Features

- 🔐 Verifies identity tokens from Privy using ECDSA public key
- 🧠 Automatically creates a user if they don't exist in the database
- 📂 Stores user by Privy ID + Wallet address
- 🧹 Clean modular backend + extensible frontend

---

## 📁 Project Structure

### Backend

```
backend/
├── internal/
│   ├── auth/          # Auth service, handler, and Privy claims
│   ├── user/          # User model, repository, and DB logic
│   └── common/utils/  # Token parser, helpers, etc.
├── main.go
```

### Frontend (Next.js + Privy)

```
frontend/
├── app/               # App router entrypoint (Next.js)
├── common/            # Shared utilities/hooks
├── components/        # UI components
├── public/            # Static files
├── .env               # Environment variables for frontend
├── next.config.ts
├── tsconfig.json
├── package.json
└── README.md
```

### Contracts (currently empty, reserved for future smart contract logic)

```
contracts/
```

---

## ⚙️ Backend Environment Variables

Create a `.env` file or export these variables manually:

```env
PRIVY_VERIFICATION_KEY="-----BEGIN PUBLIC KEY-----\n...\n-----END PUBLIC KEY-----"
PRIVY_APP_ID="your-privy-app-id"
DATABASE_URL="mongodb+srv://username:password@cluster.mongodb.net/your-db"
```


## ⚙️ Frontend Environment Variables
```env
NEXT_PUBLIC_PRIVY_APP_ID="Your_APP_ID"
NEXT_PUBLIC_PRIVY_CLIENT_ID="PRIVY_CLIENT"
NEXT_PUBLIC_SERVER_URL= "SERVER URL (this wont be needed when you deploy the server I guess)"
```

> 🔐 Make sure `PRIVY_VERIFICATION_KEY` is your **Privy's ES256 public key**, PEM-encoded and properly escaped.

---

## 🧪 Example: Auth Flow

Frontend fetches identity token via Privy:

```ts
import {useIdentityToken} from "@privy-io/react-auth";

const token = getIdentityToken()
const response = await fetch("/api/auth", {
  method: "GET",
  headers: {
    "Authorization": `Bearer ${token.identityToken}`
  }
})
```

Backend:

- Verifies token using Privy's public key
- Extracts `sub` (Privy ID)
- Extracts Wallet Address from the Identity token
- Checks if user exists in MongoDB
- If not, creates a new user
- Responds with:
  - `200 OK` – existing user
  - `201 Created` – new user created

Frontend can check `response.status === 201` to show a “Welcome new user” message.

---

## 🚀 Quick Start

1. Clone this repo  
   `git clone https://github.com/kantagara/golang-privy-mongo-boilerplate.git`

2. Set your backend and frontend `.env` files  
   Add your Privy keys and MongoDB connection string

3. Run backend  
   ```bash
   cd backend
   go run main.go
   ```

4. Run frontend  
   ```bash
   cd frontend
   npm install
   npm run dev
   ```

---

## 💡 Tips

- ✅ Don't forget to whitelist your backend URL in Privy dashboard
- ✅ Extend the `User` model to include roles, email, etc.
- ✅ Use `context.Context` for cancellations/timeouts in backend

---

## 📄 License

MIT License. Feel free to fork & build on top!
