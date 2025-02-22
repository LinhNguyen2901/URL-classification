console.log("Content script loaded");

// Add event listener to all links on the page
document.querySelectorAll('a').forEach(link => {
    link.addEventListener('mouseover', () => {
        console.log('Link href:', link.href);
    });
});