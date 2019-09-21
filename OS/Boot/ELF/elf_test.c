int printf( const char *format, ...);
int global_init_var = 80;
int global_uninit_var;

void fun1(int i) {
    printf("%d\n", i);
}

int main(){
    static int static_var = 81;
    static int static_var2;

    int a = 1;
    int b;

    fun1(static_var + static_var2 + a + b); 
}