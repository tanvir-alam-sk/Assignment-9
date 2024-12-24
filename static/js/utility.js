const votingSection = document.getElementById("voting-section");
const breedSection = document.getElementById("breed-section");
const favoriteSection = document.getElementById("favorite-section");

const toggleSection = (targetSection) => {
    votingSection.classList.remove('visible');
    breedSection.classList.remove('visible');
    favoriteSection.classList.remove('visible');

    targetSection.classList.add('visible');
}

const votingSectionOpen = () => {
    toggleSection(votingSection)
}
const breedSectionOpen = () => {
    document.getElementById("dropdown").value="abys"
    document.getElementById("caption_title").innerText="Abyssinian"
    document.getElementById("caption_description").innerText="The Abyssinian is easy to care for, and a joy to have in your home. They’re affectionate cats and love both people and other animals."
    document.getElementById("wiki").href="https://en.wikipedia.org/wiki/Abyssinian_cat"
    toggleSection(breedSection)

}

const favoriteSectionOpen = () => {
    toggleSection(favoriteSection)
    favoriteImage()
}