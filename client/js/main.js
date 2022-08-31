// Mobile Navigation

const menu = Array.from(document.querySelectorAll("[data-mobile-show]"));

const mobileToggle = document.querySelector(".mobile-toggle");

mobileToggle.addEventListener("click", () => {
  menu.forEach((i) => {
    let isToggled = i.getAttribute("data-mobile-show");

    if (isToggled === "false") {
      i.setAttribute("data-mobile-show", "true");
      mobileToggle.setAttribute("aria-expanded", "true");
    } else {
      i.setAttribute("data-mobile-show", "false");
      mobileToggle.setAttribute("aria-expanded", "false");
    }
  });
});

// Shortening

const shortenForm = document.querySelector("#main-form");
const inputElement = document.querySelector("#link");
shortenForm.addEventListener("submit", (e) => {
  e.preventDefault();
  if (isValidUrl(inputElement.value.trim())) {
    getShorteningData(inputElement.value);
  }
});

function isValidUrl(url) {
  const validUrl = Boolean(new URL(url));
  return validUrl;
}

function getShorteningData(validUrl) {
  let url = "https://api.shrtco.de/v2/shorten?url=" + validUrl;
  fetch(url)
    .then((res) => res.json())
    .then((data) => {
      processData(data.result.original_link, data.result.short_link);
    });
}

function processData(originalUrl, shortUrl) {
  const shorteningList = document.querySelector(".shortening");
  let short_link = shortUrl;
  let original_link = originalUrl;
  createShortenElement(shorteningList, original_link, short_link);
}

function createShortenElement(shortening, original_link, short_link) {
  const li = document.createElement("li");
  const p = document.createElement("p");
  p.classList.add("full-url");
  p.textContent = original_link;
  const span = document.createElement("span");
  span.classList.add("short-url");
  span.textContent = short_link;
  const button = document.createElement("button");
  button.classList.add("copy-url");
  updateBtn(button);
  li.appendChild(p);
  li.appendChild(span);
  li.appendChild(button);

  return shortening.appendChild(li);
}

function updateBtn(btn) {
  btn.textContent = "copy";

  btn.addEventListener("click", () => {
    btn.classList.add("copied");
    btn.textContent = "copied";
    navigator.clipboard.writeText(btn.previousElementSibling.textContent);
  });
}
