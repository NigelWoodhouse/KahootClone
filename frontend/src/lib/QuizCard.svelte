<script lang="ts">
    import { createEventDispatcher } from "svelte";
    import Button from "./Button.svelte";
    import GlobalLiquidGlass from "./GlobalLiquidGlass.svelte";
    import type { Quiz } from "../model/quiz";
    import { push } from "svelte-spa-router";

    const dispatch = createEventDispatcher();
    let showConfirm = $state(false);
	let itemToDelete = $state<Quiz | null>(null);

    // export let quiz: Quiz;
    let { quiz }: { quiz: Quiz } = $props();


    function host(){
        dispatch("host", quiz);
    }

    function edit() {
        push(`/edit/${quiz.id}`);
    }

    function requestDelete() {
		itemToDelete = quiz;
		showConfirm = true;
	}

	function confirmDelete() {
        dispatch("delete", quiz);
        showConfirm = false;
        itemToDelete = null;
	}

	function cancelDelete() {
		showConfirm = false;
		itemToDelete = null;
	}
</script>

{#if showConfirm}
	<dialog open class=" bg-black modal-backdrop z-10 rounded-xl text-white p-4" onclick={cancelDelete}>
        <div class="modal-content" role="presentation" onclick={(e) => e.stopPropagation()}>
			<h2>Are you sure?</h2>
			<p>You are about to delete <strong>{quiz.name}</strong>. This action cannot be undone.</p>
			<div class="actions rounded-xl sized-box flex justify-end gap-2 mt-4">
				<Button class="btn-danger" on:click={confirmDelete}>Delete</Button>
				<Button class="btn-secondary" on:click={cancelDelete}>Cancel</Button>
			</div>
		</div>
	</dialog>
{/if}

<div class="rounded-xl">
    <GlobalLiquidGlass>
        <div class="flex items-center justify-between gap-4 rounded-xl border border-white/20 p-4">
            <h1 class="text-xl font-bold text-white">{quiz.name}</h1>
            <div class="flex items-center gap-2 rounded-xl">
                <Button on:click={host}>Host</Button>
                <Button on:click={edit}>Edit</Button>
                <Button on:click={requestDelete}>Delete</Button>
            </div>
        </div>
    </GlobalLiquidGlass>
</div>

<style>

</style>