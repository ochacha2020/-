<script>
  import { fade } from "svelte/transition";
  import { fly } from "svelte/transition";

  export let drinks_type = "beer";
  const drinks = {
    beer: "🍺",
    wine: "🍷",
    cocktail: "🍸",
    tropical_drink: "🍹",
  };
  const surroundings = {
    beer: "☆",
    wine: "♡",
    cocktail: "○",
    tropical_drink: "☆",
  };

  function my_fly(node, { delay = 0, duration = 400, x = 0, y = 0 }) {
    const style = getComputedStyle(node);
    const transform = style.transform === "none" ? "" : style.transform;

    return {
      delay,
      duration,
      css: (t) => `
        transform: ${transform} translate(${(1 - t) * x}px, ${(1 - t) * y}px);
        color: hsl(
          ${~~(t * 360)},
          ${100}%,
          ${70}%
        );`,
    };
  }
</script>

<style lang="scss">
  .centerMiddle {
    margin: -200px 0 0 -300px; /*縦横半分をネガティブマージンでずらす*/
    position: absolute; /*body要素に対して絶対配置*/
    top: 50%; /*上端を中央に*/
    left: 50%; /*左端を中央に*/
    width: 640px; /*横幅*/
    height: 400px; /*縦幅*/
  }

  .drink {
    position: absolute;
    font-size: 320px;
    opacity: 1;
  }
  #drink_left {
    transform: scale(-1, 1);
    left: 70px;
  }
  #drink_right {
    left: 240px;
  }

  .surroundings {
    position: absolute;
    font-size: 7em;
    color: hsl(0, 100, 70);
  }

  $cx: 260px; //星エフェクトの中心
  $cy: 150px;
  $r: 300px; //星エフェクトの半径

  .a {
    left: $cx + $r * 1;
    bottom: $cy + $r * 0;
  }
  .b {
    left: $cx + $r * 0.5;
    bottom: $cy + $r * 0.866025403784439;
  }
  .c {
    left: $cx + $r * -0.5;
    bottom: $cy + $r * 0.866025403784439;
  }
  .d {
    left: $cx + $r * -1;
    bottom: $cy + $r * 0;
  }
  .e {
    left: $cx + $r * -0.5;
    bottom: $cy + $r * -0.866025403784439;
  }
  .f {
    left: $cx + $r * 0.5;
    bottom: $cy + $r * -0.866025403784439;
  }
</style>

<main>
  {#if drinks_type === 'beer' || drinks_type === 'wine' || drinks_type === 'cocktail' || drinks_type === 'tropical_drink'}
    <div class="centerMiddle" out:fade>
      <div class="drink" id="drink_left" in:fly={{ x: 300, duration: 100 }}>
        {drinks[drinks_type]}
      </div>
      <div class="drink" id="drink_right" in:fly={{ x: 300, duration: 100 }}>
        {drinks[drinks_type]}
      </div>

      <div class="surroundings a" in:my_fly={{ x: -100, y: 0, duration: 2000 }}>
        {surroundings[drinks_type]}
      </div>
      <div
        class="surroundings b"
        in:my_fly={{ x: -50, y: 86.6025403784439, duration: 2000 }}>
        {surroundings[drinks_type]}
      </div>
      <div
        class="surroundings c"
        in:my_fly={{ x: 50, y: 86.6025403784439, duration: 2000 }}>
        {surroundings[drinks_type]}
      </div>
      <div class="surroundings d" in:my_fly={{ x: 100, y: 0, duration: 2000 }}>
        {surroundings[drinks_type]}
      </div>
      <div
        class="surroundings e"
        in:my_fly={{ x: 50, y: -86.6025403784439, duration: 2000 }}>
        {surroundings[drinks_type]}
      </div>
      <div
        class="surroundings f"
        in:my_fly={{ x: -50, y: -86.6025403784439, duration: 2000 }}>
        {surroundings[drinks_type]}
      </div>
    </div>
    <audio id="sound" src="sound/clinking1.mp3" autoplay>
      <track kind="captions" />
    </audio>
  {/if}
</main>
