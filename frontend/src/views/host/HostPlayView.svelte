<script lang="ts">
    import Clock from "../../lib/Clock.svelte";
    import QuizChoiceCard from "../../lib/play/QuizChoiceCard.svelte";
    import { COLORS, type QuizChoice } from "../../model/quiz";
    import { type HostGame, tick, currentQuestion, gameState } from "../../service/host/host";
    import { GameState } from "../../service/net";
    import { apiService } from "../../service/api";

    function getCardColor(choice: QuizChoice, currentGameState: GameState, defaultColor: string){
        if(currentGameState != GameState.Reveal)
            return defaultColor;

        return choice.correct ? "bg-green-400" : "bg-red-400";
    }

    export let game: HostGame;
</script>

{#if $currentQuestion != null}
    <div class="min-h-screen h-screen flex flex-col">
        <div class="bg-white text-3xl border-b p-4 font-bold text-center">
            {$currentQuestion.name}
        </div>
        <div class="flex-1 flex flex-col justify-center pl-4">
            <div class="flex justify-between items-center">
                <Clock>
                    <span class="text-3xl">{$tick}</span>
                </Clock>
                <div class="w-24"></div>
            </div>
        </div>
        <div class="flex flex-wrap w-full h-96">
            {#each COLORS as color, i}
                <QuizChoiceCard color={getCardColor($currentQuestion.choices[i], $gameState, color)}>
                    <p class="pl-14">{$currentQuestion.choices[i].name}</p>
                </QuizChoiceCard>
            {/each}
        </div>
    </div>
{/if}