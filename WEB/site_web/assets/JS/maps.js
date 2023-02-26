const navLinks = document.querySelectorAll('a[href^="#"]');

// Ajouter un événement d'écouteur de clic à chaque lien de navigation
navLinks.forEach((link) => {
  link.addEventListener('click', (event) => {
    // Empêcher le comportement par défaut du clic
    event.preventDefault();
    
    // Récupérer le nom de la section cible
    const targetSectionName = link.getAttribute('href').substring(1);
    
    // Récupérer la section cible
    const targetSection = document.getElementById(targetSectionName);
    
    // Défiler en douceur jusqu'à la section cible
    targetSection.scrollIntoView({ behavior: 'smooth' });
  });
});