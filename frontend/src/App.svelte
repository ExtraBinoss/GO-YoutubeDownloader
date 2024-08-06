<script>
  import {Greet} from '../wailsjs/go/main/App.js'

  let name

  function greet() {
    Greet(name).then(result => result)
  }
</script>

<main>
  <div class="center-container">
  <div class="form" id="input">
    <input autocomplete="off" bind:value={name} class="url-input-box" id="name" placeholder="Insert YouTube URL" type="text"/>
  </div>
  <select class="format-select">
    <option value="">-- Choose Format</option>
    <option value="mp3">MP3</option>
    <option value="mp4">MP4</option>
    <option value="wav">WAV</option>
  </select>
  <button class="url-btn" on:click={greet}>Add to queue</button>
</div>
</main>

<style src="./appstyles.css"></style>
