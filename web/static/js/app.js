// Scripts de Interatividade do Portfólio Go

document.addEventListener('DOMContentLoaded', () => {
    // 1. Verificação periódica do Homelab Status
    initHomelabStatus();

    // 2. Manipulação de fechar modal com tecla ESC
    document.addEventListener('keydown', (e) => {
        if (e.key === 'Escape') {
            closeAllModals();
        }
    });
});

// Abre o modal de arquitetura do projeto
function openProjectModal(projectId) {
    const modal = document.getElementById(`modal-${projectId}`);
    if (modal) {
        modal.classList.remove('hidden');
        modal.classList.add('flex');
        document.body.style.overflow = 'hidden';
    }
}

// Fecha o modal de arquitetura
function closeProjectModal(projectId) {
    const modal = document.getElementById(`modal-${projectId}`);
    if (modal) {
        modal.classList.add('hidden');
        modal.classList.remove('flex');
        document.body.style.overflow = 'auto';
    }
}

function closeAllModals() {
    const modals = document.querySelectorAll('.project-modal');
    modals.forEach(modal => {
        modal.classList.add('hidden');
        modal.classList.remove('flex');
    });
    document.body.style.overflow = 'auto';
}

// Copiar e-mail para a área de transferência
function copyEmail(emailText, successText) {
    navigator.clipboard.writeText(emailText).then(() => {
        const btn = document.getElementById('btn-copy-email');
        const originalText = btn.innerHTML;
        btn.innerHTML = `<span class="text-emerald-400 font-medium">✓ ${successText}</span>`;
        setTimeout(() => {
            btn.innerHTML = originalText;
        }, 2500);
    }).catch(err => {
        console.error('Erro ao copiar email:', err);
    });
}

// Consulta o endpoint de status do Homelab em Golang
function initHomelabStatus() {
    fetchStatus();
    // Atualiza a cada 30 segundos
    setInterval(fetchStatus, 30000);
}

function fetchStatus() {
    fetch('/api/homelab-status')
        .then(res => res.json())
        .then(data => {
            if (data && data.services) {
                data.services.forEach((svc, index) => {
                    const badge = document.getElementById(`status-badge-${index}`);
                    if (badge) {
                        if (svc.status === 'online') {
                            badge.className = 'inline-flex items-center px-2 py-0.5 rounded-full text-xs font-medium bg-emerald-950/60 text-emerald-400 border border-emerald-800/50';
                            badge.innerHTML = `<span class="w-1.5 h-1.5 mr-1.5 rounded-full bg-emerald-400 pulse-green"></span> Online (${svc.latency})`;
                        } else {
                            badge.className = 'inline-flex items-center px-2 py-0.5 rounded-full text-xs font-medium bg-cyan-950/60 text-cyan-400 border border-cyan-800/50';
                            badge.innerHTML = `<span class="w-1.5 h-1.5 mr-1.5 rounded-full bg-cyan-400"></span> Homelab Net`;
                        }
                    }
                });
            }
        })
        .catch(err => {
            console.log('Status endpoint rodando em modo standalone local');
        });
}
