import getMessage from "@/components/messageComponent"

export default async function Home() {
  const apiResultMessage = await getMessage()

  return (
    <main className="bg-black">
      <p>MDET APP</p>
      <p>{`Message from API: ${apiResultMessage}`}</p>
    </main>
  )
}
