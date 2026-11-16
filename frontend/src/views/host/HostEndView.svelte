<script lang="ts">
    import Button from "../../lib/Button.svelte";
    import GlobalLiquidGlass from "../../lib/GlobalLiquidGlass.svelte";
    import Leaderboard from "../../lib/Leaderboard.svelte";
    import { leaderboard, type HostGame } from "../../service/host/host";

    export let game: HostGame;

    const bubbles = [
        { left: "25%", size: 100, delay: 0, duration: 25 },
        { left: "10%", size: 40, delay: 2, duration: 12 },
        { left: "70%", size: 50, delay: 4, duration: 25 },
        { left: "40%", size: 80, delay: 0, duration: 18 },
        { left: "65%", size: 60, delay: 3, duration: 25 },
        { left: "75%", size: 130, delay: 7, duration: 25 },
        { left: "35%", size: 150, delay: 15, duration: 25 },
        { left: "50%", size: 45, delay: 3, duration: 45 },
        { left: "20%", size: 35, delay: 2, duration: 35 },
        { left: "85%", size: 170, delay: 0, duration: 11 },
    ];
</script>

<div class="page-background relative isolate flex h-dvh w-full items-center justify-center overflow-hidden px-4 py-4">
    <ul class="bubbles" aria-hidden="true">
        {#each bubbles as bubble}
            <li
                style="left: {bubble.left}; width: {bubble.size}px; height: {bubble.size}px; animation-delay: {bubble.delay}s; animation-duration: {bubble.duration}s"
            ></li>
        {/each}
    </ul>
    <div class="relative z-10 flex w-full flex-col items-center">
        <h2 class="text-center text-white text-4xl font-bold sm:text-5xl">Game Ended!</h2>

        <div class="mt-4 flex w-full justify-center">
            <Leaderboard finish={true} leaderboard={$leaderboard} limit={5} />
        </div>

        <div class="mt-4">
            <Button on:click={() => game.quit()}>Quit game</Button>
        </div>
    </div>
</div>
