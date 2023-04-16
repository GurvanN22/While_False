

//Les fonctions

function Test () {

}

const Test2 = () => {

}

// Lier un element html au js avec l'id
//utilise la methode getElementById

//Ajouter un addEventListener sur un bouton pour lancer une fonction

document.getElementById('boutton').addEventListener("click" , (e) => {
    const Close = () => {
        document.getElementById('transition').classList.remove('transition')
        document.getElementById('transition').classList.add('transitionClose')
    }
    const para = document.getElementById('para')
    para.innerHTML = "bonjour"
    console.log(e.target)
    document.getElementById('hidden').style.display = "block"
    document.getElementById('boutton').style.display = "none"
    document.getElementById('transition').classList.add('transition')
    setTimeout(Close , 1000)
})

// Crée un element en js et l'ajoute à la page
//utilise la methode createElement
//utilise la methode appendChild

document.getElementById('addElement').addEventListener('click' , () => {
    const newElement = document.createElement('p')
    newElement.innerHTML = "click"
    document.getElementById('div').appendChild(newElement)
    document.getElementById("addElement").classList.add('vert')
    document.getElementById("addElement").classList.remove('rouge')
})

//Faire disparaitre les elements 
//utilise la methode replace


document.getElementById('Delete').addEventListener('click' , () => {
    document.getElementById("div").replaceChildren()
    document.getElementById("addElement").classList.remove('vert')
    document.getElementById("addElement").classList.add('rouge')
})


//modifier le css d'un element
//utilise la methode style


//Cacher un element
//utilise la methode style.display="none"

//Afficher un element
//utilise la methode style.display

//Exemple page de transition avec un bouton

//Tour de brick avec un bouton ajouter une brique
document.getElementById('addBrique').addEventListener('click' ,() => {
    const Brique = document.createElement('div')
    Brique.classList.add("Brique")
    document.getElementById('briques').appendChild(Brique)
})

