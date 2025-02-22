function storeURL(url) {
    chrome.storage.local.get({ urls: [] }, function (data) {
        let urls = data.urls;
        if (!urls.includes(url)) {  // Avoid duplicates
            urls.push(url);
            chrome.storage.local.set({ urls: urls }, function () {
                console.log("URL stored:", url);  // Debugging log
            });
        }
    });
}

function handleHover(event) {
    const link = event.target.closest("a");
    if (link) {
        console.log("Hovered URL:", link.href);  // Debugging log
        storeURL(link.href);
    }
}

document.addEventListener("mouseover", handleHover);
