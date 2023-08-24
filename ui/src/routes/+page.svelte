<script lang="ts">
  import { onMount } from "svelte";

  let ws: WebSocket
  let messages: string[] = []  
  $: {
    console.log(messages)
  }

  onMount(() => {
      ws = new WebSocket("ws://localhost:8080/ws")
      ws.onmessage = (msg) => messages = [...messages, msg.data]

      return () => {
        ws.close()
      }
  })



  let message = ''
  let username = ''
  const sendMessage = () => {
    const toSend = {username, message}
    ws.send(JSON.stringify(toSend))
    
    message = ''
  }
</script>

{#each messages as msg}
  <p>{msg}</p>
{/each}
<form on:submit|preventDefault={sendMessage}>
  <label for="username">Username</label>
  <input id="username" bind:value={username} type="text">
  <br>
  <label for="message">Message</label>
  <input id="message" bind:value={message} type="text">
  <br>
  <button type="submit">Send</button>
</form>