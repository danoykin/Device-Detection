﻿#include <stdio.h>
#include <string.h>

#include "51Degrees.h"

int main(int argc, char* argv[]) {
	char input[50000], output[50000];

    if (argc > 1) {
        char *fileName = argc > 1 ? argv[1] : NULL;
        char *requiredProperties = argc > 2 ? argv[2] : NULL;
        switch(init(fileName, requiredProperties)) {
            case DATA_SET_INIT_STATUS_INSUFFICIENT_MEMORY:
                printf("Insufficient memory to load '%s'.", argv[1]);
                break;
            case DATA_SET_INIT_STATUS_CORRUPT_DATA:
                printf("Device data file '%s' is corrupted.", argv[1]);
                break;
            case DATA_SET_INIT_STATUS_INCORRECT_VERSION:
                printf("Device data file '%s' is not correct version.", argv[1]);
                break;
            case DATA_SET_INIT_STATUS_FILE_NOT_FOUND:
                printf("Device data file '%s' not found.", argv[1]);
                break;
            default: {
                gets(input);

                while(strlen(input) > 0) {
				processDeviceCSV(getDeviceOffset(input), output, 50000);
				printf("%s", output);

				// Flush buffers.
				fflush(stdin);
				fflush(stdout);

				// Get the next useragent.
				gets(input);
			}
			destroy();
                break;
            }
        }
    } else {
        printf("Not enough arguments supplied for program to run. \n");
    }

	return 0;
}
