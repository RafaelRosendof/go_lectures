package trees

type tree struct {
	node No
	num_filhos int
	esq *tree
	dir *tree
	alt int
	}


/*
Arvore* arvore_criaVazia(); //feito

Arvore* arvore_insere(Arvore* raiz, No livro);//feito

Arvore *arv_criaNo(No livro); //feito

void arvore_libera(Arvore* raiz); //feito

bool arv_vazia(Arvore *arv); //feito

Arvore* arvore_busca(Arvore* raiz, int codigo); //feito

void arvore_buscaPorGenero(Arvore* raiz, const char genero[]); //feito

bool arvore_pertence(Arvore* raiz, int codigo);

void arvore_imprimeOrdem(Arvore* raiz); //feito
void arvore_imprimePreOrdem(Arvore* raiz);//feito
void arvore_imprimePosOrdem(Arvore* raiz);//feito

int arvore_altura(Arvore* raiz); //feito

bool removeNo(Arvore** raiz, int codigo); //feito

Arvore* carregarCSV(const char* nomeArquivo, Arvore* raiz); //TODO
*/
