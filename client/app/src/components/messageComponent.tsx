type Message = {
  message: string
}

export default async function getMessage() {
  const apiUrl = process.env.API_URL
  const res = await fetch(`${apiUrl}/ping`)
  if (!res.ok) {
    throw new Error("Failed fetch.")
  }

  const data: Message = await res.json()
  return data.message
}
