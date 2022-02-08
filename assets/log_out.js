const log_out = document.querySelector('.logout-button');
let token = localStorage.getItem('token');

log_out.addEventListener('click', () => {
    const url = window.location.origin + "/log_out";

    fetch(url, {
        method: 'POST', 
        headers: {
        'Content-Type': 'application/json',
        'Authorization': token
        },
        redirect: 'follow', 
        referrerPolicy: 'no-referrer'
    })
    localStorage.clear()
    window.location.href = window.location.origin
})

const check_token = document.querySelector('.check-button');


check_token.addEventListener('click', () => {
    const url2 = window.location.origin + "/token";

    fetch(url2, {
        method: 'POST', 
        headers: {
        'Content-Type': 'application/json',
        'Authorization': token
        },
        redirect: 'follow', 
        referrerPolicy: 'no-referrer'
    })
})