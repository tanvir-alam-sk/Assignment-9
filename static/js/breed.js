const catagoryData=[]
const selectCats=()=>{
    let breedID = document.getElementById("dropdown").value
    if(!breedID){
      breedID="abys"
    }

    fetch(`/breeds/get?breed_id=${breedID}`)
        .then(response => response.json())
        .then(data => {
          // const div=document.createElement("div")
          // div.classList.add("slide","active")

          // const img=document.createElement("img")
          // img.src=data[0].url
          // img.alt="Slide 1"

          // div.appendChild(img)
          // const slider=document.getElementById("slider")
          // slider.appendChild(div)

          // data.slice(1,4).forEach((datum,i)=>{
          //   const div=document.createElement("div")
          //   div.classList.add("slide")

          //   const img=document.createElement("img")
          //   img.src=datum.url
          //   img.alt=`Slide ${i+2}`

          //   div.appendChild(img)
          //   const slider=document.getElementById("slider")
          //   slider.appendChild(div)
          // })

         const firstImage= document.getElementById("firstImage")
         firstImage.src=data[0].url
         const secondImage= document.getElementById("secondImage")
         secondImage.src=data[1].url
         const thirdImage= document.getElementById("thirdImage")
         thirdImage.src=data[2].url
         const forthImage= document.getElementById("forthImage")
         forthImage.src=data[3].url

          selectCtagory(breedID)
        })
        .catch(error => {
          console.log(error)
        });
}


function breedCatagory(){
  const dropdown=document.getElementById("dropdown")
  
    fetch("/breeds/catagory")
        .then(response => response.json())
        .then(data => {
          (data.forEach((option)=>{
            const optionTag=document.createElement("option")
            optionTag.value=option.id
            optionTag.textContent=option.name;
            dropdown.appendChild(optionTag)
            const catagoryItem={
              id:option.id,
              name:option.name,
              description:option.description,
              wikipedia_url:option.wikipedia_url
            }
            catagoryData.push(catagoryItem)
          })) 
          
        })
        .catch(error => {
        document.getElementById("result").innerText = "Error fetching breed data.";
        });
}
breedCatagory();
selectCats();

const selectCtagory=(uid)=>{
 const selectId= catagoryData.find((data)=>data.id==uid);
 document.getElementById("caption_title").innerText=selectId.name
 document.getElementById("caption_description").innerText=selectId.description
 document.getElementById("wiki").href=selectId.wikipedia_url
}

let currentSlide = 0;
const slides = document.querySelectorAll('.slide');
const dots = document.querySelectorAll('.dot');
const dropdownBtn = document.querySelector('.dropdown-btn');
const dropdownContent = document.querySelector('.dropdown-content'); 


// Close dropdown when clicking outside
// window.addEventListener('click', (e) => {
//   if (!e.target.matches('.dropdown-btn') && !e.target.parentElement.matches('.dropdown-btn')) {
//     if (dropdownContent.classList.contains('show')) {
//       dropdownContent.classList.remove('show');
//     }
//   }
// });

function updateSlides() {
  slides.forEach(slide => slide.classList.remove('active'));
  dots.forEach(dot => dot.classList.remove('active'));
  
  slides[currentSlide].classList.add('active');
  dots[currentSlide].classList.add('active');
}

function  navigate (direction) {
  currentSlide = (currentSlide + direction + slides.length) % slides.length;
  updateSlides();
}

function changeSlide(index) {
  currentSlide = index;
  updateSlides();
  dropdownContent.classList.remove('show');
}

// Auto advance slides every 5 seconds
setInterval(() => navigate(1), 3000);