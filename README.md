# dadu_game

Merupakan game yang sangat sederhana yaitu untuk mensimulasikan permainan dadu.

Contoh command :

go run main.go --pemain=4 --dadu=5

flag --pemain untuk menentukan jumlah pemain (player)

flag --dadu untuk menentukan jumlah dadu yang akan dimainkan

Output :

Pemain = 4, Dadu = 5

==============================================================================

Giliran 1 lempar dadu

        Pemain #1 (0): 4,1,5,2,5

        Pemain #2 (0): 5,5,5,5,5

        Pemain #3 (0): 6,3,1,2,1

        Pemain #4 (0): 4,1,2,4,6

Setelah evaluasi :

        Pemain #1 (0): 4,5,2,5,1

        Pemain #2 (0): 5,5,5,5,5,1

        Pemain #3 (1): 3,2

        Pemain #4 (1): 4,2,4,1,1

==============================================================================

Giliran 2 lempar dadu

        Pemain #1 (0): 4,4,6,3,4

        Pemain #2 (0): 1,1,5,4,5,6

        Pemain #3 (1): 5,4
        
        Pemain #4 (1): 2,6,1,4,5

Setelah evaluasi :

        Pemain #1 (1): 4,4,3,4,1

        Pemain #2 (1): 5,4,5

        Pemain #3 (1): 5,4,1,1

        Pemain #4 (2): 2,4,5

==============================================================================

Giliran 3 lempar dadu

        Pemain #1 (1): 3,2,4,3,2

        Pemain #2 (1): 2,2,6

        Pemain #3 (1): 6,4,5,2

        Pemain #4 (2): 3,5,6

Setelah evaluasi :

        Pemain #1 (1): 3,2,4,3,2

        Pemain #2 (2): 2,2

        Pemain #3 (2): 4,5,2

        Pemain #4 (3): 3,5

==============================================================================

Giliran 4 lempar dadu

        Pemain #1 (1): 5,4,6,2,1

        Pemain #2 (2): 4,1

        Pemain #3 (2): 1,4,2

        Pemain #4 (3): 2,5

Setelah evaluasi :

        Pemain #1 (2): 5,4,2

        Pemain #2 (2): 4,1

        Pemain #3 (2): 4,2,1

        Pemain #4 (3): 2,5,1

==============================================================================

Giliran 5 lempar dadu

        Pemain #1 (2): 2,6,4

        Pemain #2 (2): 2,3

        Pemain #3 (2): 1,3,1

        Pemain #4 (3): 5,2,1

Setelah evaluasi :

        Pemain #1 (3): 2,4,1

        Pemain #2 (2): 2,3

        Pemain #3 (2): 3

        Pemain #4 (3): 5,2,1,1

==============================================================================

Giliran 6 lempar dadu

        Pemain #1 (3): 6,2,6

        Pemain #2 (2): 6,5

        Pemain #3 (2): 2

        Pemain #4 (3): 4,6,1,2

Setelah evaluasi :

        Pemain #1 (5): 2,1

        Pemain #2 (3): 5

        Pemain #3 (2): 2

        Pemain #4 (4): 4,2

==============================================================================

Giliran 7 lempar dadu

        Pemain #1 (5): 6,2

        Pemain #2 (3): 2

        Pemain #3 (2): 3

        Pemain #4 (4): 1,3

Setelah evaluasi :

        Pemain #1 (6): 2,1

        Pemain #2 (3): 2

        Pemain #3 (2): 3

        Pemain #4 (4): 3

==============================================================================

Giliran 8 lempar dadu

        Pemain #1 (6): 1,4

        Pemain #2 (3): 6

        Pemain #3 (2): 6

        Pemain #4 (4): 5

Setelah evaluasi :

        Pemain #1 (6): 4

        Pemain #2 (4): 1

        Pemain #3 (3): _(Berhenti bermain karena tak memiliki dadu)

        Pemain #4 (4): 5

==============================================================================

Giliran 9 lempar dadu

        Pemain #1 (6): 5

        Pemain #2 (4): 2

        Pemain #3 (3): _(Berhenti bermain karena tak memiliki dadu)

        Pemain #4 (4): 1

Setelah evaluasi :

        Pemain #1 (6): 5,1

        Pemain #2 (4): 2

        Pemain #3 (3): _(Berhenti bermain karena tak memiliki dadu)

        Pemain #4 (4): _(Berhenti bermain karena tak memiliki dadu)

==============================================================================

Giliran 10 lempar dadu

        Pemain #1 (6): 6,6

        Pemain #2 (4): 3

        Pemain #3 (3): _(Berhenti bermain karena tak memiliki dadu)

        Pemain #4 (4): _(Berhenti bermain karena tak memiliki dadu)

Setelah evaluasi :

        Pemain #1 (8): _(Berhenti bermain karena tak memiliki dadu)

        Pemain #2 (4): 3

        Pemain #3 (3): _(Berhenti bermain karena tak memiliki dadu)

        Pemain #4 (4): _(Berhenti bermain karena tak memiliki dadu)

==============================================================================

winner adalah pemain #1


