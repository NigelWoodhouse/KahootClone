<script lang="ts">
    import type { Snippet } from "svelte";
    import { LiquidGlass } from "liquid-glass-svelte";

    let { children }: { children: Snippet } = $props();

    const options = {
        mainBackgroundColor: "hsla(0, 0%, 75%, 0.1)",
        mainBlur: "1vw",
        edgeBlur: "0.5vw",
        edgeBackgroundColor: "hsla(0, 0%, 100%, 0.1)",
        edgeWidth: "1vw",
        edgeGradientWidth: "1vw",
        sheenBlur: "3vw",
        sheenBackgroundColor: "hsla(0, 0%, 100%, 0.2)",
        sheenWidth: "0.2vw",
    } as const;
</script>

<LiquidGlass class="global-liquid-glass" {options}>
    {@render children()}
</LiquidGlass>
