<script lang="ts">
    import { onMount } from "svelte";
    import QuizChoiceCard from "../../lib/play/QuizChoiceCard.svelte";
    import { COLORS } from "../../model/quiz";
    import type { PlayerGame } from "../../service/player/player";

    export let game: PlayerGame;
    let answered = false;

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

    function onClick(i: number) {
        game.answer(i);
        answered = true;
    }

    onMount(() => {
        answered = false;
    });
</script>

<div class="page-background relative isolate flex flex-wrap w-full min-h-screen overflow-hidden">
    {#if !answered}
        {#each COLORS as color, i}
            <QuizChoiceCard {color}>
                <button class="h-full w-full" on:click={() => onClick(i)}
                    >X</button
                >
            </QuizChoiceCard>
        {/each}
    {:else}
        <ul class="bubbles" aria-hidden="true">
            {#each bubbles as bubble}
                <li
                    style="left: {bubble.left}; width: {bubble.size}px; height: {bubble.size}px; animation-delay: {bubble.delay}s; animation-duration: {bubble.duration}s"
                ></li>
            {/each}
        </ul>
        <div class="relative z-10 flex w-full items-center justify-center p-8 text-center text-white">
            <p class="text-2xl font-bold">Waiting for other players to answer</p>
        </div>
    {/if}
</div>
