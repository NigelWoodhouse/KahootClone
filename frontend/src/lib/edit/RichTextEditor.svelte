<script lang="ts">
    import { createEventDispatcher, onMount } from "svelte";
    import { escapeRichText, sanitizeRichText } from "../richText";

    export let value = "";

    const dispatch = createEventDispatcher();
    let editor: HTMLDivElement;

    onMount(() => {
        editor.innerHTML = sanitizeRichText(value);
    });

    function syncValue() {
        const sanitized = sanitizeRichText(editor.innerHTML);
        if (sanitized !== editor.innerHTML) {
            editor.innerHTML = sanitized;
        }

        value = sanitized;
        dispatch("change");
    }

    function format(command: string) {
        editor.focus();
        document.execCommand(command, false);
        syncValue();
    }

    function insertClipboardContent(html: string, text: string) {
        const sanitizedHtml = html ? sanitizeRichText(html) : "";
        const plainTextHtml = escapeRichText(text).replace(/\r\n?|\n/g, "<br>");
        document.execCommand("insertHTML", false, sanitizedHtml || plainTextHtml);
        syncValue();
    }

    function handlePaste(event: ClipboardEvent) {
        event.preventDefault();
        const clipboard = event.clipboardData;
        if (clipboard) {
            insertClipboardContent(clipboard.getData("text/html"), clipboard.getData("text/plain"));
        }
    }

    function handleDrop(event: DragEvent) {
        event.preventDefault();
        const transfer = event.dataTransfer;
        if (transfer) {
            insertClipboardContent(transfer.getData("text/html"), transfer.getData("text/plain"));
        }
    }
</script>

<div class="rounded-xl border border-white/30 bg-white/10 p-2 text-white">
    <div class="mb-2 flex flex-wrap gap-1 border-b border-white/20 pb-2" role="toolbar" aria-label="Question formatting">
        <button type="button" aria-label="Bold" title="Bold" class="format-button font-bold" on:mousedown|preventDefault on:click={() => format("bold")}>B</button>
        <button type="button" aria-label="Italic" title="Italic" class="format-button italic" on:mousedown|preventDefault on:click={() => format("italic")}>I</button>
        <button type="button" aria-label="Underline" title="Underline" class="format-button underline" on:mousedown|preventDefault on:click={() => format("underline")}>U</button>
        <button type="button" aria-label="Bulleted list" title="Bulleted list" class="format-button" on:mousedown|preventDefault on:click={() => format("insertUnorderedList")}>• List</button>
        <button type="button" aria-label="Numbered list" title="Numbered list" class="format-button" on:mousedown|preventDefault on:click={() => format("insertOrderedList")}>1. List</button>
    </div>
    <div
        bind:this={editor}
        contenteditable="true"
        role="textbox"
        aria-label="Question text"
        aria-multiline="true"
        tabindex="0"
        data-placeholder="Write your question..."
        class="rich-text-editor-area rich-text-content min-h-20 max-h-48 overflow-y-auto rounded-lg bg-white p-3 text-left text-gray-900 outline-none focus:ring-2 focus:ring-purple-400"
        on:input={syncValue}
        on:paste={handlePaste}
        on:drop={handleDrop}
    ></div>
</div>

<style>
    .format-button {
        border: 1px solid rgb(255 255 255 / 20%);
        border-radius: 0.375rem;
        font-size: 0.875rem;
        padding: 0.25rem 0.625rem;
        transition: background-color 150ms ease;
    }

    .format-button:hover {
        background-color: rgb(255 255 255 / 20%);
    }

    .format-button:focus-visible {
        outline: 2px solid white;
        outline-offset: 2px;
    }
</style>
