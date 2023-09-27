type Message = {
  message: string
}

async function getData() {
  const res = await fetch("http://127.0.0.1:8080/ping", { cache: 'force-cache' })
  if (!res.ok) {
    throw new Error("Failed fetch.")
  }

  const data = await res.json()
  return data.message as Message
}

export default async function Home() {
  const apiRes = await getData()

  return (
    <main className="bg-black">
      <p>MDET APP</p>
      <p>{apiRes.message}</p>
    </main>
  )
}
