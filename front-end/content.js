const popup = document.createElement("div");
popup.className = "hover-popup";
popup.style.display = "none"; 
document.body.appendChild(popup);

function storeURL(url, link) {
    console.log("Storing URL:", url);;

    fetch("https://backend", { // Add backend API here (send url link to backend)
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify({ url: url })
    })
    .then(response => response.json())
    .then(data => {
        console.log("URL sent to backend:", data);
    })
    .catch(error => console.error("Error sending URL:", error));
}

function handleHover(event) {
    const link = event.target.closest("a");
    if (link) {
        console.log("Hovered URL:", link.href);
        storeURL(link.href, link);

        popup.textContent = "Checking ...";
        popup.style.display = "block";

        setTimeout(() => {
            popup.textContent = link.href;
        }, 500); 

        document.addEventListener("mousemove", positionPopup);
    }
}


function handleMouseLeave() {
    popup.style.display = "none";
    document.removeEventListener("mousemove", positionPopup);
}

function positionPopup(event) {
    popup.style.left = `${event.pageX + 10}px`; 
    popup.style.top = `${event.pageY - 30}px`;  
}

const style = document.createElement("style");
style.innerHTML = `

    .hover-popup {
        position: absolute;
        background: rgba(0, 0, 0, 0.8);
        color: white;
        padding: 5px 10px;
        border-radius: 5px;
        font-size: 12px;
        pointer-events: none; 
        white-space: nowrap;
        z-index: 1000;
    }
`;
document.head.appendChild(style);

document.addEventListener("mouseover", handleHover);
document.addEventListener("mouseout", handleMouseLeave);
