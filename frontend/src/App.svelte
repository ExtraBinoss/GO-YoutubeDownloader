<script>
  import {DownloadVideo} from '../wailsjs/go/main/App.js'
  import {WriteALogFile} from '../wailsjs/go/main/App.js'
  import {GetVideoAuthor, GetVideoTitle} from '../wailsjs/go/main/App.js'

  let youtube_url = '' // default url
  let format = '' // default format
  let queue_size = 0 // default queue size
  let title = '' // default title
  let author = '' // default author
  let queue_items = [];
  let checkDownloadStatus = '0';

  format.toLowerCase();
  
  async function FetchVideoAuthor(url) {
    author = await GetVideoAuthor(url);
    WriteALogFile("author fetchvideoauthor : [" + author + "]");
    if (author === '') {
      alert('Please enter a valid YouTube URL and select a format');
      throw new Error('Invalid YouTube URL');
    }
  }

  async function FetchVideoTitle(url) {
    title = await GetVideoTitle(url);
    WriteALogFile("title fetchvideotitle : [" + title + "]");
    if (title === '') {
      alert('Please enter a valid YouTube URL and select a format');
      throw new Error('Invalid YouTube URL');
    }
  }

  async function UpdateQueueSize() {
    queue_size = queue_items.length;
  }

  //let author2 = '';
  async function fn_add_to_queue() {
    console.log('youtube_url:', youtube_url);
    WriteALogFile(youtube_url);
    try {
        await FetchVideoAuthor(youtube_url);
        await FetchVideoTitle(youtube_url);
        WriteALogFile("author addtoqueue : [" + author + "]");
    } catch (error) {
        console.error('Error fetching video author:', error);
        WriteALogFile(error);
        alert('An error occurred while fetching the video author. Please try again.');
        return;
    }
      const video_id = youtube_url.split('v=')[1];
      queue_items = [
        ...queue_items,
      {
        id: video_id,
        author: author,
        title: title,
        status: 'Pending',
        format: format.toUpperCase()
      }
    ];
    UpdateQueueSize();
    processQueue();
  }

  function RemoveFromQueue(item) {
    // Delay the removal by 5 seconds // remove it completely from the list
    setTimeout(() => {
      queue_items = queue_items.filter((queue_item) => queue_item.id !== item.id);
      console.log(`Removed item: ${item}`);
      UpdateQueueSize();
    }, 2000);
  }

  async function processQueue() {
    for (let item of queue_items) {
      if (item.status === 'Pending') {
        item.status = 'Downloading...';
        checkDownloadStatus = await DownloadVideo("https://www.youtube.com/watch?v=" + item.id, item.format);
        console.log('Download Status:', checkDownloadStatus);
        if (checkDownloadStatus === '1') {
          item.status = 'Completed';
          console.log('Downloaded:', item);
          console.log("itemstatus", item.status);
          const queueBox = document.querySelector(`.queue-box[data-id="${item.id}"]`);
          if (queueBox) {
            item.status = 'Completed';
            queueBox.classList.add('completed'); // add status completed to css selector ".queue-box.completed"
            console.log('Queue Box:', queueBox);
          }
        } else {
          item.status = 'Failed';
        }
        RemoveFromQueue(item);
      }
    }
  }
  // async function fn_fetchFormat(url) {
  //   try {
  //     //push the formats to the formats array
  //     formats.push(await AvailableFormats(url));
  //     console.log('Formats:', formats);
  //   } catch (error) {
  //     console.error('Error fetching formats:', error);
  //   }
  // }

  // $: if (youtube_url) {
  //   fn_fetchFormat(youtube_url);
  // }
</script>

<main>
  <div class="center-container">
    <div class="form" id="input">
      <input autocomplete="off" bind:value={youtube_url} class="url-input-box" id="name" placeholder="Insert YouTube URL" type="text"/>
    </div>
    <select bind:value={format} class="format-select">
      <option value="">-- Choose Format</option>
      <option value="mp3">MP3</option>
      <option value="wav">WAV</option>
      <option value="mp4">MP4</option>
      <option value="mp4a">MP4a</option>
    </select>
    <button class="url-btn" on:click={fn_add_to_queue}>Add to queue</button>
  </div>
  <div class="queue-container">
    <h1 class="queue-title">Queue ({queue_size}) : </h1>
    <div class="queue-boxes">
      {#each queue_items as item}
      <div class="queue-box" data-id={item.id}>
        <img src={`https://i.ytimg.com/vi/${item.id}/hqdefault.jpg`} alt="Thumbnail" class="thumbnail">
          <div class="queue-details">
            <h2 class="queue-title">{item.title || 'Loading...'}</h2>
            <p class="queue-author">Author: {item.author}</p>
            <p class="queue-format">Format: {item.format}</p>
            <p class="queue-status">Status: {item.status}</p>
          </div>
        </div>
      {/each}
    </div>
  </div>
</main>

<style src="./appstyles.css"></style>


<!-- TO DO -->
<!-- List all formats by parsing it using 144, 360, 480p etc... -->
<!-- Choose path to download -->