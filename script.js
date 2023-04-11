var NSFW_enabled = false;

function createPhraseBox(endpoint) {
    fetch(endpoint)
      .then(response => response.json())
      .then(data => {
        const phraseBox = document.querySelector('.phrase-box');
        const phraseContent = document.querySelector('.phrase-content');
  
        phraseContent.innerHTML = data.content;
  
        phraseBox.classList.add('revealed');
      });
  }
  
  document.querySelector('#NSFW_disable_button').addEventListener('click', () => {
    NSFW_enabled = false;
    const phraseBox = document.querySelector('.phrase-box');
    phraseBox.style.boxShadow = "0 0 10px rgb(17, 236, 218)";
    var nsfwElement = document.querySelector('.nsfw-check');
    nsfwElement.textContent = 'NSFW OFF';
  });
  
  document.querySelector('#NSFW_enable_button').addEventListener('click', () => {
    NSFW_enabled = true;
    const phraseBox = document.querySelector('.phrase-box');
    phraseBox.style.boxShadow = "0 0 10px rgb(17, 236, 218)";
    var nsfwElement = document.querySelector('.nsfw-check');
    nsfwElement.textContent = 'NSFW ON';
  });
  
  document.querySelector('#truth_button').addEventListener('click', () => {
    const phraseBox = document.querySelector('.phrase-box');
    phraseBox.style.boxShadow = "0 0 15px rgba(8, 46, 219, 0.836)";

    if (NSFW_enabled) {
      createPhraseBox('http://tod.clowzy0.com:8080/get_ANY_truth');
    }

    if (!NSFW_enabled) {
      createPhraseBox('http://tod.clowzy0.com:8080/get_SFW_truth');
    }
  });
  
  document.querySelector('#dare_button').addEventListener('click', () => {
    const phraseBox = document.querySelector('.phrase-box');
    phraseBox.style.boxShadow = "0 0 15px rgba(82, 14, 170, 0.836)";

    if (NSFW_enabled) {
      createPhraseBox('http://tod.clowzy0.com:8080/get_ANY_dare');
    }

    if (!NSFW_enabled) {
      createPhraseBox('http://tod.clowzy0.com:8080/get_SFW_dare');
    }
  });

  /*
  document.querySelector('#truth_button').addEventListener('click', () => {
    const phraseBox = document.querySelector('.phrase-box');
    phraseBox.style.boxShadow = "0 0 15px rgba(8, 46, 219, 0.836)";

    if (NSFW_enabled) {
      createPhraseBox('http://tod.clowzy0.com:8080/get_ANY_truth');
    }

    if (!NSFW_enabled) {
      createPhraseBox('http://tod.clowzy0.com:8080/get_SFW_truth');
    }
  });
  
  document.querySelector('#dare_button').addEventListener('click', () => {
    const phraseBox = document.querySelector('.phrase-box');
    phraseBox.style.boxShadow = "0 0 15px rgba(82, 14, 170, 0.836)";

    if (NSFW_enabled) {
      createPhraseBox('http://tod.clowzy0.com:8080/get_ANY_dare');
    }

    if (!NSFW_enabled) {
      createPhraseBox('http://tod.clowzy0.com:8080/get_SFW_dare');
    }
  });
  */
