<script>
  import Mogumogu from "./Mogumogu.svelte";
  import Cheer_effect from "./Cheer_effect.svelte";
  import Emoticon from "./Emoticon.svelte";
  import { callVisonAPI } from "./api/SendRequest";
  import { getEmocon } from "./api/Emotion";
  import Fab, { Label, Icon } from "@smui/fab";
  export let onBack = () => {};

  $: feeling = "";
  let my_mode = "normal";
  let modes = ["normal", "flower", "cat", "lovers"];
  let drinks_type = "beer";
  let cheers_now = false;
  let food_type = "USA";
  let mogumogu_now = false;

  let peer;
  let localStream;
  let flag = true;
  let theirID = "";

  // カメラ映像取得
  navigator.mediaDevices
    .getUserMedia({ video: true, audio: true })
    .then((stream) => {
      // 成功時にvideo要素にカメラ映像をセットし、再生
      const videoElm = document.getElementById("my-video");
      videoElm.srcObject = stream;
      videoElm.play();
      // 着信時に相手にカメラ映像を返せるように、グローバル変数に保存しておく
      localStream = stream;
      if (flag) {
        flag = false;
        api_call();
      }
    })
    .catch((error) => {
      // 失敗時にはエラーログを出力
      console.error("mediaDevice.getUserMedia() error:", error);
      return;
    });

  peer = new Peer({
    key: "",
    debug: 3,
  });

  peer.on("open", () => {
    document.getElementById("my-id").textContent = peer.id;
  });

  //着信処理
  peer.on("call", (mediaConnection) => {
    mediaConnection.answer(localStream);
    setEventListener(mediaConnection);
    theirID = mediaConnection.remoteId;
  });

  peer.on("close", () => {
    alert("通信が切断しました。");
  });

  peer.on("error", (err) => {
    alert(err.message);
  });

  // データを送信する

  function sendEmotion(feeling) {
    const emo_messages = document.getElementById("js-emo-messages");
    const dataConnection = peer.connect(theirID);
    peer.on("connection", (dataConnection) => {
      const data = {
        name: "SkyWay", //dataConnection.id,
        msg: feeling,
      };
      dataConnection.send(data);
    });

    peer.connect("peerID").on("data", ({ name, msg }) => {
      emo_messages.textContent += `${name}: ${msg}`;
      // => 'SkyWay: Hello, World!'
    });
  }

  // イベントリスナを設置する関数
  const setEventListener = (mediaConnection) => {
    mediaConnection.on("stream", (stream) => {
      // video要素にカメラ映像をセットして再生
      const videoElm = document.getElementById("their-video");
      videoElm.srcObject = stream;
      videoElm.play();
    });
  };

  // 発信処理
  const makeCall = () => {
    theirID = document.getElementById("their-id").value;
    const mediaConnection = peer.call(theirID, localStream);
    setEventListener(mediaConnection);
  };

  function sleep(ms) {
    return new Promise((resolve) => setTimeout(resolve, ms));
  }

  const captureMovie = () => {
    var video = document.getElementById("my-video");
    const canvasSizeX = 400;
    const canvasSizeY = 300;
    var canvas = document.getElementById("draw");
    var context = canvas.getContext("2d");
    context.drawImage(video, 0, 0, canvasSizeX, canvasSizeY);
    call_vision_API(canvas.toDataURL());
  };

  async function call_vision_API(baseImg) {
    const emo_messages = document.getElementById("js-emo-messages");
    emo_messages.textContent += `${peer.id}:call API\n`;
    var json_res = await callVisonAPI(baseImg);
    feeling = getEmocon(json_res);
    sendEmotion(feeling);
  }

  async function api_call() {
    for (let step = 0; step < 5; step++) {
      await sleep(5000);
      captureMovie();
    }
  }

  const onBackHome = () => {
    videoEl.srcObject.getTracks().forEach(function (track) {
      track.stop();
    });
    onBack();
  };
</script>

<style>
  main {
    text-align: center;
    padding: 1em;
    max-width: 240px;
    margin: 0 auto;
  }

  @media (min-width: 640px) {
    main {
      max-width: none;
    }
  }

  .relative {
    position: relative;
    width: 400px;
    margin: 0 auto;
  }
</style>

<svelte:head />
<main>
  <div class="relative">
    <video id="my-video" width="400px" autoplay muted playsinline />
    <!--リアクションが映像の中に現れる-->
    <Emoticon {feeling} mode={my_mode} />
  </div>

  <video id="their-video" width="400px" autoplay muted playsinline />
  <p id="my-id" />
  <textarea id="their-id" />
  <button on:click={makeCall}>発信</button>

  <label> <input type="checkbox" bind:checked={cheers_now} /> cheers! </label>
  {#if cheers_now}
    <Cheer_effect {drinks_type} />
  {/if}

  <label>
    <input type="checkbox" bind:checked={mogumogu_now} /> mogumogu now!
  </label>
  {#if mogumogu_now}
    <Mogumogu {food_type} />
  {/if}

  <!-- 非表示ゾーン hidden-->
  <pre class="emo-messages" id="js-emo-messages" />
  <canvas id="draw" width="400px" height="300px" hidden />

  <!--モードを入力-->
  <label>
    <input type="radio" bind:group={my_mode} value={'normal'} /> normal
  </label>
  <label>
    <input type="radio" bind:group={my_mode} value={'flower'} /> flower
  </label>
  <label> <input type="radio" bind:group={my_mode} value={'cat'} /> cat </label>
  <label>
    <input type="radio" bind:group={my_mode} value={'lovers'} /> lovers
  </label>

  <!--飲み物の種類を入力-->
  <label>
    <input type="radio" bind:group={drinks_type} value={'beer'} /> beer
  </label>
  <label>
    <input type="radio" bind:group={drinks_type} value={'wine'} /> wine
  </label>
  <label>
    <input type="radio" bind:group={drinks_type} value={'cocktail'} /> cocktail
  </label>
  <label>
    <input type="radio" bind:group={drinks_type} value={'tropical_drink'} /> tropical_drink
  </label>

  <!--食べ物の種類を入力-->
  <label>
    <input type="radio" bind:group={food_type} value={'USA'} /> USA
  </label>
  <label>
    <input type="radio" bind:group={food_type} value={'bread'} /> bread
  </label>
  <label>
    <input type="radio" bind:group={food_type} value={'Japan'} /> Japan
  </label>
  <label>
    <input type="radio" bind:group={food_type} value={'sweets'} /> sweets
  </label>

  <Fab color="primary" on:click={onBack}>
    <Icon class="material-icons appbutton">home</Icon>
  </Fab>
</main>
