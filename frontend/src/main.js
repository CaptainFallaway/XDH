import {GetHello} from '../wailsjs/go/app/App';

window.render = function() {
    // Call App.GetHello()
    try {
        GetHello()
            .then((result) => {
                // Update result with data back from App.GetHello()
                document.querySelector('#models').innerHTML = result;
            })
            .catch((err) => {
                console.error(err);
            });
    } catch (err) {
        console.error(err);
    }
}