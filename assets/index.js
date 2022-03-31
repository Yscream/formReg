const firstName = document.querySelector('.name');
const lastName = document.querySelector('.lastName');
const email = document.querySelector('.email');
const password = document.querySelector('.password');
const buttonSignUp = document.querySelector('.button');

buttonSignUp.addEventListener('click', signUp);


function signUp(){
    const url = window.location.origin + "/user";
    
    const formValue = {
        Name: firstName.value,
        LastName: lastName.value,
        Email: email.value,
        Password: password.value,
    }

    const response = fetch(url, {
        method: 'POST', 
        mode: 'no-cors', 
        cache: 'no-cache',
        credentials: 'same-origin', 
        headers: {
        'Content-Type': 'application/json'
        },
        redirect: 'follow', 
        referrerPolicy: 'no-referrer',
        body: JSON.stringify(formValue) 
    })
    .then((response) => {
            return response.json();})
    .then((data) => { 
        if(data.length === 0){
            window.location.href = 'link.html';
        }
        if (data.length > 0){
            for (let obj of data) {
                if (obj.FieldName === 'Name') {
                    document.querySelector('#error1').innerHTML = obj.MessageErr;
                    oninput(firstName, '#error1')
                }
                if (obj.FieldName === 'LastName') {
                    document.querySelector('#error2').innerHTML = obj.MessageErr;
                    oninput(lastName, '#error2')
                }
                if (obj.FieldName === 'Email') {
                    document.querySelector('#error3').innerHTML = obj.MessageErr;
                    oninput(email, '#error3')
                }
                if (obj.FieldName === 'Password') {
                    document.querySelector('#error4').innerHTML = obj.MessageErr;
                    oninput(password, '#error4')
                }
            }
        }
    }); 
}

function oninput(field, row) {
    field.oninput = () => {
        document.querySelector(row).innerHTML = '';
    }
}
