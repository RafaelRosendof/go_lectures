package redblack


/*
// Funções básicas
tree* arv_criaVazia();                                 // Cria uma árvore vazia
No* arv_criaNo(No* pai, No node, Color color);         // Cria um novo nó
void arv_libera(tree* arv);                            // Libera a árvore
bool arv_vazia(tree* arv);                             // Verifica se a árvore está vazia

// Funções específicas para Red-Black Tree
tree* arv_insereRB(tree* arv, No node);                // Insere um nó na árvore
void corrigeInsercao(No* no);                          // Corrige a árvore após inserção
No* rotacaoEsquerda(tree* arv, No* no);                // Rotação à esquerda
No* rotacaoDireita(tree* arv, No* no);                 // Rotação à direita

bool arv_removeRB(tree* arv, int score);               // Remove um nó da árvore
void corrigeRemocao(No* no);                           // Corrige a árvore após remoção

// Funções auxiliares
void imprimeOrdem(No* raiz);                           // Imprime os nós em ordem
void imprimePreOrdem(No* raiz);                        // Imprime os nós em pré-ordem
No* arv_busca(No* raiz, int score);                    // Busca um nó com base no score
int arv_altura(No* raiz);                              // Calcula a altura da árvore
int max(int a, int b);                                 // Retorna o máximo entre dois valores

ver uma lib de go para imprimir em formato de desenho

*/
type Color string

const (
	Red   Color = "Red"
	Black Color = "Black"
)

type Node struct {
	score int
	cor   Color
	esq   *Node
	dir   *Node
	pai   *Node
}

type tree struct {
	raiz *Node
}
