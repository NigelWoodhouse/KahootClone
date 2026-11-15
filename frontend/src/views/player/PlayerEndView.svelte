<script lang="ts">
    import { leaderboard, totalPoints, playerName } from "../../service/player/player";
    import GlobalLiquidGlass from "../../lib/GlobalLiquidGlass.svelte";


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

<div class="page-background relative isolate min-h-screen w-full overflow-hidden px-6 py-12 text-white">
    <ul class="bubbles" aria-hidden="true">
        {#each bubbles as bubble}
            <li
                style="left: {bubble.left}; width: {bubble.size}px; height: {bubble.size}px; animation-delay: {bubble.delay}s; animation-duration: {bubble.duration}s"
            ></li>
        {/each}
    </ul>

    <main class="relative flex flex-col items-center text-center">
        <h1 class="text-4xl font-bold sm:text-5xl">Game Ended!</h1>

            <div class="flex rounded-xl p-4 flex-col max-w-2xl">
                <GlobalLiquidGlass>

                <p class="rounded-xl px-8 text-2xl font-bold shadow-lg">
                    {$playerName} <br>
                    Your final score: {$totalPoints} points
                </p>
                </GlobalLiquidGlass>

            </div>
    </main>
</div>
