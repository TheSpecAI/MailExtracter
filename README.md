# MailExtracter

**An automated Gmail extraction service by Spec.AI**

### 🚀 Overview

MailExtracter is a service that retrieves emails from a Gmail account (with user permissions) and provides an API to fetch them securely.

### 🔧 Features

- **Automated email retrieval** every 5 minutes.
- **Secure API route** to fetch extracted emails (requires a secret key).
- **OAuth-based authentication** to access Gmail data.

### 📌 Setup

#### 1️⃣ Clone the Repository

```sh
git clone https://github.com/TheSpecAI/MailExtracter.git
cd MailExtracter
```

#### 2️⃣ Install Dependencies

```sh
go mod tidy
```

#### 3️⃣ Configure Environment Variables

- Set up a **Google Cloud project** and enable **Gmail API**.
- Obtain **OAuth credentials** (Client ID & Secret).
- Check `.env.example` for required environment variables and create a `.env` file:
- Add **credentials.json** to root folder

#### 4️⃣ Run the Service using hot-reloader

```sh
air
```

### 🔑 API Usage

#### ➤ Get Emails

```http
POST /mail
```

##### **Request Body:**

```json
{
  "secret_key": "yourkey"
}
```

##### **Response:**

Returns a list of retrieved emails if the correct secret key is provided.

### 🛠 Future Improvements

- Support for multiple email providers.
- Webhook-based real-time email fetching.

### 💡 Contributing

Feel free to submit issues or pull requests!
