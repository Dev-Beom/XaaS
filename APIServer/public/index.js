const generateBtn = document.querySelector('.generateBtn');
const idInput = document.querySelector('#id')
const descriptionInput = document.querySelector('#description')

const generateNode = async () => {
    await fetch("http://localhost:5000/api/node", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify({
            id: idInput.value,
            description: descriptionInput.value
        }),
    })
}

generateBtn.addEventListener('click', generateNode);
