<script>
  import RoomVideo from "./RoomVideo.svelte";
  import { onMount } from "svelte";
  import Emoticon from "./Emoticon.svelte";
  import Cheer_effect from "./Cheer_effect.svelte";
  import Mogumogu from "./Mogumogu.svelte";
  import { callVisonAPI } from "./api/SendRequest";
  import { getEmocon } from "./api/Emotion";
  import P2P from "./P2P.svelte";
  import Fab, { Label, Icon } from "@smui/fab";
  export let onBack = () => {};
  import Radio from "@smui/radio";
  import FormField from "@smui/form-field";
  let imageSrc = "image/emoapiIcon.jpg";
  let setting = false;
  export let roomName = "";
  let messages = "";
  let isStartCalling = false;

  $: feeling = "";
  let mode = "normal";
  let room;
  let my_mode = "normal";
  let modes = ["normal", "flower", "cat", "lovers"];
  let drinks_type = "beer";
  let cheers_setup_now = false;
  let cheers_now = false;
  let member = 1;
  let food_type = "USA";
  let food_type0 = "USA";
  let food_type2 = "USA";
  let mogumogu_setup_now = false;
  let mogumogu_now = false;
  let mogumogu_now0 = false;
  let mogumogu_now2 = false;
  let mogumogu_set_now0 = false;
  let mogumogu_set_now2 = false;
  const getRoomModeByHash = () => (location.hash === "#sfu" ? "sfu" : "mesh");

  const positions = {
    1: {},
    2: {},
  };
  const Peer = window.Peer;
  const user = {
    id: 0,
    name: "sample name",
    muted: true,
    camOn: true,
    srcObject: null,
    playsInline: true,
    autoplay: true,
    width: 400,
  };
  let owner = Object.create(user);

  let remoteUsers = [];

  var sendEmo = function () {
    return "null";
  };
  let flag = true;

  function user1(id, feeling, mode) {
    this.id = id;
    this.feeling = feeling;
    this.mode = mode;
  }

  function user2(id, feeling, mode) {
    this.id = id;
    this.feeling = feeling;
    this.mode = mode;
  }

  var user1_state = new user1("", "", "normal");
  var user2_state = new user2("", "", "normal");

  // eslint-disable-next-line require-atomic-updates
  const peer = (window.peer = new Peer({
    key: "",
    debug: 3,
  }));

  function updateVideo(video, _user) {
    video.muted = _user.muted;
    video.srcObject = _user.srcObject;
    video.playsInline = _user.playsInline;
    video.width = _user.width;
    video.height - _user.width;
  }

  onMount(async () => {
    const joinTrigger = document.getElementById("js-join-trigger");
    const leaveTrigger = document.getElementById("js-leave-trigger");
    //const remoteVideos = document.getElementById("js-remote-streams");
    const remoteTwoVideos = document.getElementById("js-remote-two-streams");
    const remoteThreeVideos = document.getElementById(
      "js-remote-three-streams"
    );
    const video0 = document.getElementById("video0");
    const video1 = document.getElementById("video1");
    const video2 = document.getElementById("video2");
    video0.width = 0;
    video2.width = 0;
    const roomId = document.getElementById("js-room-id");
    const sdkSrc = document.querySelector("script[src*=skyway]");
    api_call();

    const localStream = await navigator.mediaDevices
      .getUserMedia({
        audio: true,
        video: true,
      })
      .catch(console.error);

    owner.id = peer.id;
    owner.width = 400;
    owner.srcObject = localStream;

    // Register join handler
    joinTrigger.addEventListener("click", () => {
      if (peer.open && roomName != "") {
        isStartCalling = true;

        const room = peer.joinRoom(roomName, {
          mode: getRoomModeByHash(),
          stream: localStream,
        });

        room.on("peerJoin", (peerId) => {
          messages += `=== ${peerId} joined ===\n`;
        });

        // Render remote stream for new peer join in the room
        room.on("stream", async (stream) => {
          member += 1;

          if (member === 2) {
            video0.width = 400;
            video0.height = 400;
            remoteTwoVideos.srcObject = stream;
            remoteTwoVideos.playsInline = true;
            remoteTwoVideos.width = 400;
            remoteTwoVideos.height = 400;
            remoteTwoVideos.id = stream.peerId;
            remoteTwoVideos.muted = false;
            remoteTwoVideos.setAttribute("data-peer-id", stream.peerId);
            user1_state = new user1(stream.peerId, "", "");
            await remoteTwoVideos.play().catch(console.error);
          } else if (member === 3) {
            video2.width = 400;
            video2.height = 400;
            remoteThreeVideos.srcObject = stream;
            remoteThreeVideos.playsInline = true;
            remoteThreeVideos.width = 400;
            remoteThreeVideos.height = 400;
            remoteThreeVideos.id = stream.peerId;
            remoteThreeVideos.setAttribute("data-peer-id", stream.peerId);
            remoteThreeVideos.muted = false;
            user2_state = new user2(stream.peerId, "", "");
            await remoteThreeVideos.play().catch(console.error);
          }
        });

        // このコードできたデータの受け取りを行う
        room.on("data", ({ data, src }) => {
          if (data == "cheer:beer") {
            drinks_type = "beer";
            cheer();
            return;
          } else if (data == "cheer:wine") {
            drinks_type = "wine";
            cheer();
            return;
          } else if (data == "cheer:cocktail") {
            drinks_type = "cocktail";
            cheer();
            return;
          } else if (data == "cheer:tropical_drink") {
            drinks_type = "tropical_drink";
            cheer();
            return;
          }

          if (user1_state.id == src) {
            if (data.length > 7) {
              let mogu_data = data.substr(0, 7);
              let _foodtype = data.substr(7, data.length - 7);
              if (mogu_data == "mogu-on") {
                mogumogu_set_now0 = true;
                food_type0 = _foodtype;
                mogumogu0();
                return;
              } else if (mogu_data == "mogu-of") {
                mogumogu_set_now0 = false;
                return;
              } else if (mogu_data == "mute") {
                return;
              } else if (mogu_data == "unmute") {
                return;
              }
            }
            if (data.startsWith("Emotion")) {
              let emo_data = data.split(";");
              if (emo_data.length == 3) {
                user1_state.feeling = emo_data[1];
                user1_state.mode = emo_data[2];
              }
            }
          }

          if (user2_state.id == src) {
            if (data.length > 7) {
              let mogu_data = data.substr(0, 7);
              let _foodtype = data.substr(7, data.length - 7);
              if (mogu_data == "mogu-on") {
                mogumogu_set_now2 = true;
                food_type2 = _foodtype;
                mogumogu2();
                return;
              } else if (mogu_data == "mogu-of") {
                mogumogu_set_now2 = false;
                return;
              } else if (mogu_data == "mute") {
                return;
              } else if (mogu_data == "unmute") {
                return;
              }
            }

            if (data.startsWith("Emotion")) {
              let emo_data = data.split(";");
              if (emo_data.length == 3) {
                user2_state.feeling = emo_data[1];
                user2_state.mode = emo_data[2];
              }
            }
          }
        });

        // for closing room members
        room.on("peerLeave", (peerId) => {
          member -= 1;
          let remoteVideo = remoteTwoVideos.querySelector(
            `[data-peer-id="${peerId}"]`
          );
          if (remoteVideo == null) {
            remoteVideo = remoteThreeVideos.querySelector(
              `[data-peer-id="${peerId}"]`
            );
          }
          remoteVideo.srcObject.getTracks().forEach((track) => track.stop());
          remoteVideo.srcObject = null;
          remoteVideo.remove();

          // 残っている絵文字を消す
          if (peerID == user1_state.id) {
            user1_state.feeling = "";
            user1_state.mode = "";
          }

          if (peerID == user2_state.id) {
            user2_state.feeling = "";
            user2_state.mode = "";
          }
        });

        // for closing myself
        room.once("close", () => {
          member -= 1;
          //sendTrigger.removeEventListener("click", onClickSend);
          messages += "== You left ===\n";
          Array.from(remoteTwoVideos).forEach((remoteVideo) => {
            remoteVideo.srcObject.getTracks().forEach((track) => track.stop());
            remoteVideo.srcObject = null;
            remoteVideo.remove();
          });
          Array.from(remoteThreeVideos).forEach((remoteVideo) => {
            remoteVideo.srcObject.getTracks().forEach((track) => track.stop());
            remoteVideo.srcObject = null;
            remoteVideo.remove();
          });
          onBack();
        });

        // API呼び出し関連
        sendEmo = function sendEmotion(feeling) {
          const emo_messages = document.getElementById("js-emo-messages");
          // 表情を他の人に伝える
          const emotion = `Emotion;${feeling};${my_mode}`;
          room.send(emotion);
        };

        leaveTrigger.addEventListener("click", () => room.close(), {
          once: true,
        });
      }
    });

    peer.on("error", console.error);
  });

  // API呼び出し関連
  function sleep(ms) {
    return new Promise((resolve) => setTimeout(resolve, ms));
  }

  // APIを呼び出して送る
  async function call_vision_API(baseImg) {
    var json_res = await callVisonAPI(baseImg);
    feeling = getEmocon(json_res);
    sendEmo(feeling);
  }
  const captureMovie = () => {
    var video = document.getElementById("js-local-stream");
    const canvasSizeX = 400;
    const canvasSizeY = 300;
    var canvas = document.getElementById("draw");
    var context = canvas.getContext("2d");
    context.drawImage(video, 0, 0, canvasSizeX, canvasSizeY);
    call_vision_API(canvas.toDataURL());
  };

  async function api_call() {
    for (let step = 0; step < 50; step++) {
      await sleep(6000);
      captureMovie();
    }
  }

  function onBackHome() {
    onBack();
    room.close();
  }

  const onMogu = () => {
    mogumogu_setup_now = !mogumogu_setup_now;
    if (mogumogu_setup_now) {
      sendEmo("mogu-on" + food_type);
      mogumogu_now = true;
      mogumogu();
    } else {
      sendEmo("mogu-of" + food_type);
    }
  };

  $: {
    if (mogumogu_now) {
      sendEmo("mogu-on" + food_type);
    }
  }

  function mogumogu2() {
    if (mogumogu_set_now2) {
      mogumogu_now2 = true;
      setTimeout(() => {
        mogumogu_now2 = false;
        setTimeout(() => {
          mogumogu2();
        }, 1000);
      }, 10000);
    } else {
      mogumogu_now2 = false;
    }
  }

  function mogumogu0() {
    if (mogumogu_set_now0) {
      mogumogu_now0 = true;
      setTimeout(() => {
        mogumogu_now0 = false;
        setTimeout(() => {
          mogumogu0();
        }, 1000);
      }, 10000);
    } else {
      mogumogu_now0 = false;
    }
  }

  function mogumogu() {
    if (mogumogu_setup_now) {
      mogumogu_now = true;
      setTimeout(() => {
        mogumogu_now = false;
        setTimeout(() => {
          mogumogu();
        }, 1000);
      }, 10000);
    } else {
      mogumogu_now = false;
    }
  }

  function onCheerSet() {
    cheers_setup_now = !cheers_setup_now;
    cheers_now = false;
  }

  function onCheer() {
    sendEmo("cheer:" + drinks_type);
    cheer();
  }

  function cheer() {
    cheers_now = true;
    setTimeout(() => {
      cheers_now = false;
    }, 2000);
  }

  function onOpenSetting() {
    setting = !setting;
  }
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

  .container-max {
    width: 100%;
  }

  @media (min-width: 640px) {
    main {
      max-width: none;
    }
  }

  .container-video {
    display: flex;
    flex-direction: row;
    height: 360pt;
  }

  .container-fab {
    display: flex-end;
    position: fixed;
    bottom: 16pt;
    right: 16pt;
  }

  .fab-position {
    margin-top: 5pt;
    bottom: 0;
    right: 0;
  }

  .kanpai-position {
    margin-top: 15pt;
  }

  .kanpai-contain {
    display: flex;
    flex-direction: column;
  }

  .videoSize {
    width: 400pt;
    height: 400pt;
  }

  .flex-container {
    display: flex;
    justify-content: flex-end;
  }

  .flex-container > :first-child {
    margin-right: auto;
  }

  .relative {
    position: relative;
    margin-top: 16pt;
    width: 400pt;
  }
</style>

<div class="flex-container">
  <main class="container-max">
    <div class="container-max">
      <div>
        <div class="container-video">
          <div id="video0" class="relative">
            <Emoticon feeling={user1_state.feeling} mode={user1_state.mode} />
            <video
              class="video-mask trim"
              id="js-remote-two-streams"
              muted
              playsInline />
            {#if mogumogu_now0}
              <Mogumogu {food_type0} />
            {/if}
          </div>
          <div id="video1" class="relative videoSize">
            <RoomVideo
              id={owner.id}
              name={owner.name}
              muted={owner.muted}
              camOn={owner.camOn}
              srcObject={owner.srcObject}
              imgeSrc={imageSrc}
              playsInline={owner.playsInline}
              width={owner.width}
              start="false" />
            <!--リアクションが映像の中に現れる-->
            <Emoticon {feeling} mode={my_mode} />
            {#if mogumogu_now}
              <Mogumogu {food_type} />
            {/if}
          </div>

          <!-- 非表示ゾーン hidden-->
          <div id="video2" class="relative">
            <Emoticon feeling={user2_state.feeling} mode={user2_state.mode} />
            <video
              class="video-mask trim"
              id="js-remote-three-streams"
              muted
              playsInline />
            {#if mogumogu_now2}
              <Mogumogu {food_type2} />
            {/if}
          </div>
        </div>
      </div>

      <pre class="messages" id="js-messages" />
      {#if cheers_now}
        <Cheer_effect {drinks_type} />
      {/if}

      <!-- TODO -->

      {#if !isStartCalling}
        <Fab id="js-join-trigger" color="secondary">
          <Icon class="material-icons">call</Icon>
        </Fab>
      {/if}

      <!--モードを入力-->
      {#if !cheers_setup_now && !mogumogu_setup_now && isStartCalling}
        <p>今の気分は？</p>
        <FormField>
          <Radio bind:group={my_mode} value="normal" />
          <span slot="label">Normal</span>
        </FormField>
        <FormField>
          <Radio bind:group={my_mode} value="flower" />
          <span slot="label">Flower</span>
        </FormField>
        <FormField>
          <Radio bind:group={my_mode} value="cat" />
          <span slot="label">Cat</span>
        </FormField>
        <FormField>
          <Radio bind:group={my_mode} value="lovers" />
          <span slot="label">Lovers</span>
        </FormField>
      {/if}

      <!--飲み物の種類を入力-->
      {#if cheers_setup_now}
        <div class="kanpai-contain">
          <p>何で乾杯する？</p>
          <div>
            <FormField>
              <Radio bind:group={drinks_type} value="beer" />
              <span slot="label">Beer</span>
            </FormField>
            <FormField>
              <Radio bind:group={drinks_type} value="wine" />
              <span slot="label">Wine</span>
            </FormField>
            <FormField>
              <Radio bind:group={drinks_type} value="cocktail" />
              <span slot="label">Cocktail</span>
            </FormField>
            <FormField>
              <Radio bind:group={drinks_type} value="tropical_drink" />
              <span slot="label">Tropical drink</span>
            </FormField>
          </div>
          <div class="kanpai-position">
            <Fab on:click={onCheer} extended>
              <Label>かんぱ〜い！</Label>
            </Fab>
          </div>
        </div>
      {/if}

      <!--食べ物の種類を入力-->
      {#if mogumogu_setup_now}
        <p>もぐもぐタイム</p>
        <FormField>
          <Radio bind:group={food_type} value="USA" />
          <span slot="label">USA</span>
        </FormField>
        <FormField>
          <Radio bind:group={food_type} value="bread" />
          <span slot="label">bread</span>
        </FormField>
        <FormField>
          <Radio bind:group={food_type} value="Japan" />
          <span slot="label">Japan</span>
        </FormField>
        <FormField>
          <Radio bind:group={food_type} value="sweets" />
          <span slot="label">sweets</span>
        </FormField>
      {/if}
    </div>
  </main>
  <div class="container-fab">
    <div class="fab-position">
      {#if !cheers_setup_now && setting}
        <Fab color="secondary" on:click={onCheerSet}>
          <Icon class="material-icons">wine_bar</Icon>
        </Fab>
      {/if}
      {#if cheers_setup_now && setting}
        <Fab color="primary" on:click={onCheerSet}>
          <Icon class="material-icons">wine_bar</Icon>
        </Fab>
      {/if}
    </div>
    <div class="fab-position">
      {#if !mogumogu_setup_now && setting}
        <Fab color="secondary" on:click={onMogu}>
          <Icon class="material-icons">fastfood</Icon>
        </Fab>
      {/if}
      {#if mogumogu_setup_now && setting}
        <Fab color="primary" on:click={onMogu}>
          <Icon class="material-icons">fastfood</Icon>
        </Fab>
      {/if}
    </div>
    <div class="fab-position">
      {#if setting}
        <Fab color="secondary" on:click={onBackHome}>
          <Icon class="material-icons">call_end</Icon>
        </Fab>
      {/if}
    </div>
    <div class="fab-position">
      {#if setting}
        <Fab on:click={onOpenSetting}>
          <Icon class="material-icons">close</Icon>
        </Fab>
      {/if}
      {#if !setting}
        <Fab on:click={onOpenSetting}>
          <Icon class="material-icons">settings</Icon>
        </Fab>
      {/if}
    </div>
  </div>
</div>

<pre class="emo-messages videoSize" id="js-emo-messages" />
<canvas id="draw" width="400px" height="300px" hidden />
