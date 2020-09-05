<script>
  import { onMount } from "svelte";
  import "./theme/button.scss";
  export let id;
  export let name = "sample";
  export let muted = true;
  export let camOn = false;
  export let srcObject = null;
  export let imgeSrc = null;
  export let playsInline = true;
  export let width = 400;
  export let start = false;
  let loading = true;
  let localVideo;

  $: {
    localVideo = document.getElementById("js-local-stream");
    if (localVideo != null) {
      localVideo.muted = muted;
      localVideo.srcObject = srcObject;
      localVideo.playsInline = playsInline;
      localVideo.width = width;
      localVideo.height = width;
      start = true;
    }
  }

  $: {
    if (localVideo != null) {
      localVideo.srcObject = srcObject;
      start = true;
    }
  }

  $: changePlayState(start).then((value) => {
    loading = false;
  });

  async function changePlayState(isStart) {
    if (isStart) {
      if (localVideo != null && !localVideo.playing) {
        await localVideo.play().catch(console.error);
      }
    } else {
      if (!loading) {
        localVideo.srcObject.getTracks().forEach((track) => track.stop());
        localVideo.srcObject = null;
        localVideo.remove();
      }
    }
  }
  window.onload = function () {};

  onMount(() => {
    if (srcObject != null) {
      localVideo.srcObject = srcObject;
    }
    if (loading && srcObject != null) {
      start = true;
      loading = false;
    }
  });
</script>

<style>
  main {
    text-align: center;
    padding: 1em;
    max-width: 240px;
    margin: 0 auto;
  }

  .video-mask {
    border-radius: 50% 50% 50% 50%;
    object-fit: cover;
  }

  @media (min-width: 640px) {
    main {
      max-width: none;
    }
  }
</style>

<main>
  {#if !loading && camOn}
    <video class="video-mask trim" id="js-local-stream" muted playsInline />
  {/if}
  {#if !loading && !camOn}
    <p>loading...</p>
    <img src={imgeSrc} {width} height={width} alt="user-icon" />
  {/if}
  {#if loading}
    <p>{name}はCamOffです</p>
  {/if}
</main>
