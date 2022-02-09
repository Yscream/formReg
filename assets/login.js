const loginEmail = document.querySelector('.logEmail')
const loginPassword = document.querySelector('.logPassword')
const buttonLogin = document.querySelector('#login')

function logIn(){
    const url = window.location.origin + "/log";
    
    const logForm = {
        Email: loginEmail.value,
        Password: loginPassword.value,
    }

    fetch(url, {
        method: 'POST', 
        mode: 'no-cors', 
        cache: 'no-cache',
        credentials: 'same-origin', 
        headers: {
        'Content-Type': 'application/json'
        },
        redirect: 'follow', 
        referrerPolicy: 'no-referrer',
        body: JSON.stringify(logForm) 
    })
    .then((response) => {
        return response.json();})
    .then((data) => { 
        if(typeof(data) === 'string'){
            localStorage.setItem('token', data)
            window.location.href = 'after_log.html';
        }
        if(data.length > 0){
            for (let obj of data) {
                if (obj.FieldName === 'Email') {
                    document.querySelector('#error3').innerHTML = obj.MessageErr;
                    oninput(loginEmail, '#error3')
                }
                if (obj.FieldName === 'Password') {
                    document.querySelector('#error4').innerHTML = obj.MessageErr;
                    oninput(loginPassword, '#error4')
                }
            }
        }
    })
}

function oninput(field, row) {
    field.oninput = () => {
        document.querySelector(row).innerHTML = '';
    }
}
buttonLogin.addEventListener('click', logIn);