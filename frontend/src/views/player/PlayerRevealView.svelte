<script>
    import { points, totalPoints } from "../../service/player/player";

    $: correct = $points > 0;
</script>

<div
    class="min-h-screen text-white w-full {correct
        ? 'bg-green-500'
        : 'bg-red-600'} flex justify-center items-center"
>
    {#if correct}
    <div class="text-center">
        <h2 class="text-3xl font-bold">Correct!</h2>
        <p class="text-2xl">+ {$points} points</p>
        <p class="mt-2 text-lg">Total score: {$totalPoints}</p>
        </div>
    {:else}
        <div class="text-center">
            <h2 class="text-3xl">Incorrect!</h2>
            <p class="mt-2 text-lg">0 points this question</p>
            <p class="mt-2 text-lg">Total score: {$totalPoints}</p>
        </div>
    {/if}
</div>
