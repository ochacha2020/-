<script>
  import "./theme/button.scss";
  import P2P from "./P2P.svelte";
  import Room from "./Room.svelte";
  import IconButton from "@smui/icon-button";
  import Textfield from "@smui/textfield";
  import Button, { Group, GroupItem, Label, Icon } from "@smui/button";
  import TopAppBar, {
    Row,
    Section,
    Title,
    FixedAdjust,
    DenseFixedAdjust,
    ProminentFixedAdjust,
    ShortFixedAdjust,
  } from "@smui/top-app-bar";
  import { space } from "svelte/internal";
  export let toggleText;
  let imageSrc = "image/emoapiIcon.jpg";
  let collapsed = true;
  // init -> p2p or room -> init -> ...という感じで表示させる．
  let screenState = "init";
  let toggleState = true;
  let ready = false;
  let loading = false;
  let variant = "short";
  let title = "taitoru";
  let roomName = "";

  const initializeSkyWay = () => {
    ready = true;
  };

  let onBack = () => {
    screenState = "init";
  };

  let goP2P = () => {
    screenState = "p2p";
  };

  let goRoom = () => {
    if (roomName != "") {
      screenState = "room";
    }
  };

  let toggle = () => {
    if (toggleState) {
      toggleText = "Roomを動かしてます";
    } else {
      toggleText = "P2Pを動かしてます";
    }
    toggleState = !toggleState;
  };
</script>

<style>
  main {
    text-align: center;
    padding: 1em;
    max-width: 240px;
    margin: 0 auto;
  }

  .center {
    margin: 0 auto;
    margin-top: 10%;
  }

  .container0 {
    display: flex;
    flex-direction: column;
  }

  .text-field-container {
    width: 200pt;
    margin: auto;
    margin-top: -10pt;
    margin-bottom: 20pt;
  }

  @media (min-width: 640px) {
    main {
      max-width: none;
    }
  }
</style>

<TopAppBar {variant}>
  <Row>
    <Section>
      <IconButton class="material-icons">menu</IconButton>
    </Section>
  </Row>
</TopAppBar>

<main>
  {#if loading}
    <!-- TODO: 回転ビール -->
  {/if}
  {#if screenState == 'init'}
    <!-- htmlわからん誰か真ん中に持ってこさせて．． -->
    <div class="container0">
      <img
        class="center"
        src={imageSrc}
        width="200px"
        height="200px"
        alt="icon" />
      <div class="text-field-container">
        <Textfield bind:value={roomName} label="ROOM NAME" />
      </div>
    </div>
    <div>
      <Button variant="unelevated" class="button-shaped-round" on:click={goP2P}>
        <Icon class="material-icons">favorite</Icon>
        <Label>2人で会話</Label>
      </Button>
      <Button
        variant="unelevated"
        class="button-shaped-round"
        on:click={goRoom}>
        <Icon class="material-icons">groups</Icon>2人以上で会話
      </Button>
    </div>
  {/if}
  {#if screenState == 'p2p' && ready}
    <P2P {onBack} />
  {/if}

  {#if screenState == 'room' && ready}
    <Room {onBack} {roomName} />
  {/if}
  <script
    src="https://cdn.webrtc.ecl.ntt.com/skyway-latest.js"
    on:load={initializeSkyWay}>
  </script>
</main>
