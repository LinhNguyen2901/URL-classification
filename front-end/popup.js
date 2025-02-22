document.addEventListener("DOMContentLoaded", function () {
    const urlList = document.getElementById("url-list");
    const clearButton = document.getElementById("clear-urls");

    // Load saved URLs from Chrome storage
    chrome.storage.local.get({ urls: [] }, function (data) {
        data.urls.forEach(url => {
            let li = document.createElement("li");
            let a = document.createElement("a");
            a.href = url;
            a.textContent = url;
            a.target = "_blank";  // Open link in a new tab
            li.appendChild(a);
            urlList.appendChild(li);
        });
    });

    // Clear saved URLs
    clearButton.addEventListener("click", function () {
        chrome.storage.local.set({ urls: [] }, function () {
            urlList.innerHTML = "";  // Clear UI
        });
    });
});
