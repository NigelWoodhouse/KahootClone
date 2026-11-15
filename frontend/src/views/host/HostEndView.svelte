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

<div class="page-background relative isolate min-h-screen w-full overflow-hidden px-6 py-12">
    <ul class="bubbles" aria-hidden="true">
        {#each bubbles as bubble}
            <li
                style="left: {bubble.left}; width: {bubble.size}px; height: {bubble.size}px; animation-delay: {bubble.delay}s; animation-duration: {bubble.duration}s"
            ></li>
        {/each}
    </ul>
    <div class="relative z-10 flex flex-col items-center">
        <h2 class="text-center text-white text-5xl font-bold">Game Ended!</h2>

        <div class="flex flex-wrap justify-center gap-2 mt-10 w-full">
            <Leaderboard finish={true} leaderboard={$leaderboard} />
        </div>

        <div class="mt-8">
            <Button on:click={() => game.quit()}>Quit game</Button>
        </div>
    </div>
</div>
