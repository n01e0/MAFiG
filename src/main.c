#include <stdio.h>
#include <string.h>
#include <stdlib.h>

int Init(FILE* dest){
    fprintf(dest,"#include <stdio.h>\n");
    fprintf(dest,"#include <stdlib.h>\n\n");
    fprintf(dest,"#define MEM_MAX 32767\n");
    fprintf(dest,"int input() {\n");
    fprintf(dest,"\tint c = getchar();\n");
    fprintf(dest,"\tif(c == EOF){\n");
    fprintf(dest,"\t\texit(0);\n");
    fprintf(dest,"\t}\n");
    fprintf(dest,"\treturn c;\n");
    fprintf(dest,"}\n");
    fprintf(dest,"\n");
    fprintf(dest,"int main() {\n");
    fprintf(dest,"\tchar* ptr = (char*)malloc(MEM_MAX);\n");
    return 0;
}

int trans(FILE* src, FILE* dest){
    char c;
    while(c != EOF){
        switch(c = fgetc(src)) {
            case '+':
                fprintf(dest,"\t(*ptr)++;\n");
                break;
            case '-':
                fprintf(dest,"\t(*ptr)--;\n");
                break;
            case '>':
                fprintf(dest,"\tptr++;\n");
                break;
            case '<':
                fprintf(dest,"\tptr--;\n");
                break;
            case ',':
                fprintf(dest,"\t*ptr = input();\n");
                break;
            case '.':
                fprintf(dest,"\tputchar(*ptr);\n"); 
                break;
            case '[':
                fprintf(dest,"\twhile(*ptr){\n"); 
                break;
            case ']':
                fprintf(dest,"\t};\n");
                break;
        }
    }

    fprintf(dest,"}\n");
}


int main(int argc, char **argv){
    if(argc < 2){
        printf("Usage: %s [brainfucked file]\n", argv[0]);
        return 1;
    }

    int PathLen = strlen(argv[1]) + 3;
    char DestPath[PathLen];
    snprintf(DestPath, PathLen, "%s%s", argv[1], ".c");

    FILE *dest;
    dest = fopen(DestPath, "w+");
    Init(dest);

    FILE *src;
    src = fopen(argv[1], "r+");

    trans(src, dest);
    
    fclose(src);
    fclose(dest);
    return 0;
}
