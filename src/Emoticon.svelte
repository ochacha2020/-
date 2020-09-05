<script>
  import { fade } from "svelte/transition";
  import { elasticOut } from "svelte/easing";

  export let mode = "normal";
  export let feeling;

  const perfectly_joy_center = {
    normal: "😆",
    flower: "😆",
    cat: "😸",
    lovers: "🥰",
  };
  const perfectly_joy_surroundings = {
    normal: "⭐",
    flower: "🌼",
    cat: "🐟",
    lovers: "🌺",
  };
  const soso_joy_center = {
    normal: "😊",
    flower: "😊",
    cat: "😺",
    lovers: "😊",
  };
  const soso_joy_surroundings = {
    normal: "🍀",
    flower: "🌸",
    cat: "🍀",
    lovers: "💕",
  };
  const angry_center = { normal: "😠", flower: "😠", cat: "😾", lovers: "🙃" };
  const angry_surroundings = {
    normal: "💭",
    flower: "💭",
    cat: "💭",
    lovers: "💭",
  };
  const sad_center = { normal: "😢", flower: "😢", cat: "😿", lovers: "🥺" };
  const sad_surroundings = {
    normal: "💧",
    flower: "💧",
    cat: "🌀",
    lovers: "💙",
  };
  const surprised_center = {
    normal: "😲",
    flower: "😲",
    cat: "🙀",
    lovers: "😲",
  };
  const surprised_surroundings = {
    normal: "❗",
    flower: "❗",
    cat: "❗",
    lovers: "❗",
  };

  function spin(node, { duration, delay }) {
    return {
      duration,
      delay,

      css: (t) => {
        const eased = elasticOut(t);
        return `
          transform: scale(${eased}) rotate(${eased * 1080}deg);
        `;
      },
    };
  }

  function spinAndFadeOut(node, { duration, delay }) {
    return {
      duration,
      delay,

      css: (t, u) => {
        const eased = elasticOut(t);
        return `
          transform: scale(${eased}) rotate(${eased * 1080}deg);
          opacity: ${u}
        `;
      },
    };
  }

  function fadeIn(node, { duration, delay }) {
    return {
      delay,
      duration,

      css: (t) => `opacity: ${t}`,
    };
  }

  function big(node, { delay, duration = 4000 }) {
    return {
      delay,
      duration,

      css: (t) => {
        const eased = elasticOut(t);
        return `
          transform: scale(${eased * 2});
          opacity: 1;
        `;
      },
    };
  }
</script>

<style lang="scss">
  $cx: 27px; //エフェクトの中心
  $cy: 67px;
  $r: 70px; //エフェクトの半径

  .center {
    position: absolute;
    font-size: 80px;
    right: $cx - 25;
    top: $cy - 25;
  }
  .surroundings {
    position: absolute;
    font-size: 34px;
    opacity: 0;
  }

  .a {
    right: $cx + $r * 1;
    top: $cy + $r * 0;
  }
  .b {
    right: $cx + $r * 0.5;
    top: $cy + $r * 0.866025403784439;
  }
  .c {
    right: $cx + $r * -0.5;
    top: $cy + $r * 0.866025403784439;
  }
  .d {
    right: $cx + $r * -1;
    top: $cy + $r * 0;
  }
  .e {
    right: $cx + $r * -0.5;
    top: $cy + $r * -0.866025403784439;
  }
  .f {
    right: $cx + $r * 0.5;
    top: $cy + $r * -0.866025403784439;
  }
</style>

<main>
  {#if feeling === 'perfectly_joy' && (mode === 'normal' || mode === 'flower' || mode === 'cat' || mode === 'lovers')}
    <div class="absolute" out:fade>
      <div class="center" in:spin={{ duration: 8000, delay: 0 }}>
        {perfectly_joy_center[mode]}
      </div>
      <div
        class="surroundings a"
        in:spinAndFadeOut={{ duration: 4000, delay: 1000 }}>
        {perfectly_joy_surroundings[mode]}
      </div>
      <div
        class="surroundings b"
        in:spinAndFadeOut={{ duration: 4000, delay: 1000 }}>
        {perfectly_joy_surroundings[mode]}
      </div>
      <div
        class="surroundings c"
        in:spinAndFadeOut={{ duration: 4000, delay: 1000 }}>
        {perfectly_joy_surroundings[mode]}
      </div>
      <div
        class="surroundings d"
        in:spinAndFadeOut={{ duration: 4000, delay: 1000 }}>
        {perfectly_joy_surroundings[mode]}
      </div>
      <div
        class="surroundings e"
        in:spinAndFadeOut={{ duration: 4000, delay: 1000 }}>
        {perfectly_joy_surroundings[mode]}
      </div>
      <div
        class="surroundings f"
        in:spinAndFadeOut={{ duration: 4000, delay: 1000 }}>
        {perfectly_joy_surroundings[mode]}
      </div>
    </div>
  {/if}

  {#if feeling === 'soso_joy' && (mode === 'normal' || mode === 'flower' || mode === 'cat' || mode === 'lovers')}
    <div class="absolute" out:fade>
      <div class="center" in:fade={{ duration: 1000 }}>
        {soso_joy_center[mode]}
      </div>
      <div class="surroundings a" in:fadeIn={{ duration: 2000, delay: 500 }}>
        {soso_joy_surroundings[mode]}
      </div>
      <div class="surroundings b" in:fadeIn={{ duration: 2000, delay: 2500 }}>
        {soso_joy_surroundings[mode]}
      </div>
      <div class="surroundings c" in:fadeIn={{ duration: 2000, delay: 1000 }}>
        {soso_joy_surroundings[mode]}
      </div>
      <div class="surroundings d" in:fadeIn={{ duration: 2000, delay: 2000 }}>
        {soso_joy_surroundings[mode]}
      </div>
      <div class="surroundings e" in:fadeIn={{ duration: 2000, delay: 3000 }}>
        {soso_joy_surroundings[mode]}
      </div>
      <div class="surroundings f" in:fadeIn={{ duration: 2000, delay: 1500 }}>
        {soso_joy_surroundings[mode]}
      </div>
    </div>
  {/if}

  {#if feeling === 'angry' && (mode === 'normal' || mode === 'flower' || mode === 'cat' || mode === 'lovers')}
    <div class="absolute" out:fade>
      <div class="center" in:fade={{ duration: 1000 }}>
        {angry_center[mode]}
      </div>
      <div class="surroundings a" in:fadeIn={{ duration: 1000, delay: 500 }}>
        {angry_surroundings[mode]}
      </div>
      <div class="surroundings c" in:fadeIn={{ duration: 1000, delay: 1000 }}>
        {angry_surroundings[mode]}
      </div>
      <div class="surroundings b" in:fadeIn={{ duration: 1000, delay: 2500 }}>
        {angry_surroundings[mode]}
      </div>
      <div class="surroundings d" in:fadeIn={{ duration: 1000, delay: 2000 }}>
        {angry_surroundings[mode]}
      </div>
      <div class="surroundings e" in:fadeIn={{ duration: 1000, delay: 3000 }}>
        {angry_surroundings[mode]}
      </div>
      <div class="surroundings f" in:fadeIn={{ duration: 1000, delay: 1500 }}>
        {angry_surroundings[mode]}
      </div>
    </div>
  {/if}

  {#if feeling === 'sad' && (mode === 'normal' || mode === 'flower' || mode === 'cat' || mode === 'lovers')}
    <div class="absolute" out:fade>
      <div class="center" in:fade={{ duration: 1000 }}>{sad_center[mode]}</div>
      <div class="surroundings a" in:fadeIn={{ duration: 1000, delay: 500 }}>
        {sad_surroundings[mode]}
      </div>
      <div class="surroundings c" in:fadeIn={{ duration: 1000, delay: 1000 }}>
        {sad_surroundings[mode]}
      </div>
      <div class="surroundings b" in:fadeIn={{ duration: 1000, delay: 2500 }}>
        {sad_surroundings[mode]}
      </div>
      <div class="surroundings d" in:fadeIn={{ duration: 1000, delay: 2000 }}>
        {sad_surroundings[mode]}
      </div>
      <div class="surroundings e" in:fadeIn={{ duration: 1000, delay: 3000 }}>
        {sad_surroundings[mode]}
      </div>
      <div class="surroundings f" in:fadeIn={{ duration: 1000, delay: 1500 }}>
        {sad_surroundings[mode]}
      </div>
    </div>
  {/if}

  {#if feeling === 'surprised' && (mode === 'normal' || mode === 'flower' || mode === 'cat' || mode === 'lovers')}
    <div class="absolute" out:fade>
      <div class="center" in:fade={{ duration: 1000 }}>
        {surprised_center[mode]}
      </div>
      <div class="surroundings a" in:big={{ delay: 1000 }}>
        {surprised_surroundings[mode]}
      </div>
      <div class="surroundings c" in:big={{ delay: 1000 }}>
        {surprised_surroundings[mode]}
      </div>
      <div class="surroundings b" in:big={{ delay: 1000 }}>
        {surprised_surroundings[mode]}
      </div>
      <div class="surroundings d" in:big={{ delay: 1000 }}>
        {surprised_surroundings[mode]}
      </div>
      <div class="surroundings e" in:big={{ delay: 1000 }}>
        {surprised_surroundings[mode]}
      </div>
      <div class="surroundings f" in:big={{ delay: 1000 }}>
        {surprised_surroundings[mode]}
      </div>
    </div>
  {/if}
</main>
