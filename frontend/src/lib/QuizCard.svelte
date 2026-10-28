<script lang="ts">
    import { createEventDispatcher } from "svelte";
    import Button from "./Button.svelte";
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
	<!-- svelte-ignore a11y_click_events_have_key_events -->
	<!-- svelte-ignore a11y_no_noninteractive_element_interactions -->
	<dialog open class="modal-backdrop" onclick={cancelDelete}>
		<!-- Stop propagation so clicking inside the box doesn't close it -->
        <div class="modal-content" role="presentation" onclick={(e) => e.stopPropagation()}>
			<h2>Are you sure?</h2>
			<p>You are about to delete <strong>{quiz.name}</strong>. This action cannot be undone.</p>
			<div class="actions">
				<Button class="btn-danger" on:click={confirmDelete}>Delete</Button>
				<Button class="btn-secondary" on:click={cancelDelete}>Cancel</Button>
			</div>
		</div>
	</dialog>
{/if}

<div class="flex justify-between items-center bg-white border p-4 rounded-xl">
    <p>{quiz.name}</p>
    <div class="flex gap-2 items-center">
        <Button on:click={host}>Host</Button>
        <Button on:click={edit}>Edit</Button>
        <Button on:click={requestDelete}>Delete</Button>
    </div>
</div>