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
    <input autocomplete="off" bind:value={name} class="url-input-box" id="name" type="text"/>
  </div>
  <button class="url-btn" on:click={greet}>Greet</button>
</div>
</main>

<style src="./appstyles.css"></style>
