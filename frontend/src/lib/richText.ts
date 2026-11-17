const ALLOWED_TAGS = new Set([
    "B", "BLOCKQUOTE", "BR", "CODE", "DEL", "DIV", "EM",
    "H1", "H2", "H3", "H4", "H5", "H6", "I", "LI",
    "OL", "P", "S", "STRONG", "SUB", "SUP", "U", "UL",
]);

const DROP_CONTENT_TAGS = new Set([
    "EMBED", "IFRAME", "MATH", "OBJECT", "SCRIPT", "STYLE",
    "SVG", "TEMPLATE",
]);

const BLOCK_TAGS = new Set([
    "BLOCKQUOTE", "DIV", "H1", "H2", "H3", "H4", "H5",
    "H6", "LI", "OL", "P", "UL",
]);

function escapeHtml(value: string): string {
    return value
        .replace(/&/g, "&amp;")
        .replace(/</g, "&lt;")
        .replace(/>/g, "&gt;")
        .replace(/"/g, "&quot;")
        .replace(/'/g, "&#39;");
}

function serializeNode(node: Node): string {
    if (node.nodeType === 3) {
        return escapeHtml(node.nodeValue ?? "");
    }

    if (node.nodeType !== 1) {
        return Array.from(node.childNodes).map(serializeNode).join("");
    }

    const element = node as Element;
    if (DROP_CONTENT_TAGS.has(element.tagName)) {
        return "";
    }

    const content = Array.from(element.childNodes).map(serializeNode).join("");
    if (element.tagName === "BR") {
        return "<br>";
    }

    if (ALLOWED_TAGS.has(element.tagName)) {
        const tag = element.tagName.toLowerCase();
        return `<${tag}>${content}</${tag}>`;
    }

    const styles = element.getAttribute("style")?.toLowerCase() ?? "";
    let formattedContent = content;
    if (/(?:^|;)\s*text-decoration(?:-line)?\s*:[^;]*\bunderline\b/.test(styles)) {
        formattedContent = `<u>${formattedContent}</u>`;
    }
    if (/(?:^|;)\s*font-style\s*:\s*italic\b/.test(styles)) {
        formattedContent = `<em>${formattedContent}</em>`;
    }
    if (/(?:^|;)\s*font-weight\s*:\s*(?:bold|[6-9]00)\b/.test(styles)) {
        formattedContent = `<strong>${formattedContent}</strong>`;
    }

    return formattedContent;
}

export function sanitizeRichText(value: string): string {
    if (typeof DOMParser === "undefined") {
        return escapeHtml(value);
    }

    const document = new DOMParser().parseFromString(value, "text/html");
    return Array.from(document.body.childNodes).map(serializeNode).join("");
}

function extractPlainText(node: Node): string {
    if (node.nodeType === 3) {
        return node.nodeValue ?? "";
    }

    const content = Array.from(node.childNodes).map(extractPlainText).join("");
    if (node.nodeType !== 1) {
        return content;
    }

    const element = node as Element;
    return element.tagName === "BR" || BLOCK_TAGS.has(element.tagName)
        ? `${content} `
        : content;
}

export function richTextToPlainText(value: string): string {
    if (typeof DOMParser === "undefined") {
        return value;
    }

    const document = new DOMParser().parseFromString(sanitizeRichText(value), "text/html");
    return extractPlainText(document.body).replace(/\s+/g, " ").trim();
}

export function escapeRichText(value: string): string {
    return escapeHtml(value);
}
