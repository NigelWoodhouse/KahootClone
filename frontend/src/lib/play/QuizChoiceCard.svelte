<script lang="ts">
    export let color: string;
    export let fixedHeight = false;
</script>

<div class="w-1/2 {fixedHeight ? 'h-1/2' : ''} text-white text-center font-bold flex items-center justify-center text-3xl {color}">
    <slot />
</div>