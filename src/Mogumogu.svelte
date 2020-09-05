<script>
  import { fade } from "svelte/transition";
  import { fly } from "svelte/transition";

  export let food_type = "USA";
  const food1 = { USA: "🍖", bread: "🥐", Japan: "🍣", sweets: "🍦" };
  const food2 = { USA: "🍔", bread: "🥖", Japan: "🍶", sweets: "🍧" };
  const food3 = { USA: "🍟", bread: "🥨", Japan: "🍢", sweets: "🍩" };
  const food4 = { USA: "🌮", bread: "🥯", Japan: "🌰", sweets: "🍨" };
  const food5 = { USA: "🧀", bread: "🥞", Japan: "🍵", sweets: "🍪" };
  const food6 = { USA: "🍗", bread: "🍞", Japan: "🍱", sweets: "🍰" };
  const food7 = { USA: "🍕", bread: "🧇", Japan: "🍘", sweets: "🍭" };
</script>

<style>
  .pukapuka {
    position: absolute;
    font-size: 80px;
  }

  .a {
    left: 110px;
    bottom: 280px;
  }
  .b {
    left: 250px;
    bottom: 280px;
  }
  .c {
    left: 60px;
    bottom: 240px;
  }
  .d {
    left: 300px;
    bottom: 240px;
  }
  .e {
    left: 180px;
    bottom: 240px;
  }
  .f {
    left: 110px;
    bottom: 210px;
  }
  .g {
    left: 250px;
    bottom: 210px;
  }
</style>

<main>
  {#if food_type === 'USA' || food_type === 'bread' || food_type === 'Japan' || food_type === 'sweets'}
    <div class="pukapuka a" in:fly={{ y: 280, duration: 15000 }} out:fade>
      {food1[food_type]}
    </div>
    <div
      class="pukapuka b"
      in:fly={{ y: 280, duration: 15000, delay: 400 }}
      out:fade>
      {food2[food_type]}
    </div>
    <div
      class="pukapuka c"
      in:fly={{ y: 220, duration: 13000, delay: 800 }}
      out:fade>
      {food3[food_type]}
    </div>
    <div
      class="pukapuka d"
      in:fly={{ y: 220, duration: 13000, delay: 1200 }}
      out:fade>
      {food4[food_type]}
    </div>
    <div
      class="pukapuka e"
      in:fly={{ y: 220, duration: 13000, delay: 1600 }}
      out:fade>
      {food5[food_type]}
    </div>
    <div
      class="pukapuka f"
      in:fly={{ y: 210, duration: 11000, delay: 2000 }}
      out:fade>
      {food6[food_type]}
    </div>
    <div
      class="pukapuka g"
      in:fly={{ y: 210, duration: 11000, delay: 2400 }}
      out:fade>
      {food7[food_type]}
    </div>
  {/if}
</main>
