package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Struktur Book untuk menyimpan data buku
type Book struct {
	ID     int
	Title  string
	Author string
	Year   int
}

var books []Book // Slice untuk menyimpan daftar buku
var nextID int   // Variabel untuk menyimpan ID buku berikutnya

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		// Menampilkan menu perpustakaan
		fmt.Println("\nLibrary Menu:")
		fmt.Println("1. Add Book")
		fmt.Println("2. Edit Book")
		fmt.Println("3. Delete Book")
		fmt.Println("4. List Books")
		fmt.Println("5. Exit")
		fmt.Print("Select an option: ")
		option, _ := reader.ReadString('\n')
		option = strings.TrimSpace(option)

		// Menjalankan fungsi berdasarkan opsi yang dipilih
		switch option {
		case "1":
			addBook(reader) // Menambahkan buku
		case "2":
			editBook(reader) // Mengedit buku
		case "3":
			deleteBook(reader) // Menghapus buku
		case "4":
			listBooks() // Menampilkan daftar buku
		case "5":
			fmt.Println("Exiting...")
			return // Keluar dari program
		default:
			fmt.Println("Invalid option, please try again.")
		}
	}
}

// Fungsi untuk menambahkan buku baru
func addBook(reader *bufio.Reader) {
	fmt.Print("Enter book title: ")
	title, _ := reader.ReadString('\n')
	title = strings.TrimSpace(title)

	fmt.Print("Enter book author: ")
	author, _ := reader.ReadString('\n')
	author = strings.TrimSpace(author)

	fmt.Print("Enter book year: ")
	yearStr, _ := reader.ReadString('\n')
	yearStr = strings.TrimSpace(yearStr)
	year, err := strconv.Atoi(yearStr) // Mengkonversi tahun ke tipe integer
	if err != nil {
		fmt.Println("Invalid year format, book not added.")
		return
	}

	nextID++ // Meningkatkan ID buku berikutnya
	book := Book{
		ID:     nextID,
		Title:  title,
		Author: author,
		Year:   year,
	}
	books = append(books, book) // Menambahkan buku ke dalam slice books
	fmt.Println("Book added successfully.")
}

// Fungsi untuk mengedit buku yang ada
func editBook(reader *bufio.Reader) {
	fmt.Print("Enter book ID to edit: ")
	idStr, _ := reader.ReadString('\n')
	idStr = strings.TrimSpace(idStr)
	id, err := strconv.Atoi(idStr) // Mengkonversi ID ke tipe integer
	if err != nil {
		fmt.Println("Invalid ID format.")
		return
	}

	// Mencari buku berdasarkan ID
	for i, book := range books {
		if book.ID == id {
			fmt.Print("Enter new book title (leave blank to keep current): ")
			title, _ := reader.ReadString('\n')
			title = strings.TrimSpace(title)
			if title != "" {
				book.Title = title
			}

			fmt.Print("Enter new book author (leave blank to keep current): ")
			author, _ := reader.ReadString('\n')
			author = strings.TrimSpace(author)
			if author != "" {
				book.Author = author
			}

			fmt.Print("Enter new book year (leave blank to keep current): ")
			yearStr, _ := reader.ReadString('\n')
			yearStr = strings.TrimSpace(yearStr)
			if yearStr != "" {
				year, err := strconv.Atoi(yearStr) // Mengkonversi tahun ke tipe integer
				if err == nil {
					book.Year = year
				} else {
					fmt.Println("Invalid year format, keeping current year.")
				}
			}

			books[i] = book // Memperbarui buku dalam slice books
			fmt.Println("Book updated successfully.")
			return
		}
	}

	fmt.Println("Book not found.")
}

// Fungsi untuk menghapus buku
func deleteBook(reader *bufio.Reader) {
	fmt.Print("Enter book ID to delete: ")
	idStr, _ := reader.ReadString('\n')
	idStr = strings.TrimSpace(idStr)
	id, err := strconv.Atoi(idStr) // Mengkonversi ID ke tipe integer
	if err != nil {
		fmt.Println("Invalid ID format.")
		return
	}

	// Mencari dan menghapus buku berdasarkan ID
	for i, book := range books {
		if book.ID == id {
			books = append(books[:i], books[i+1:]...) // Menghapus buku dari slice books
			fmt.Println("Book deleted successfully.")
			return
		}
	}

	fmt.Println("Book not found.")
}

// Fungsi untuk menampilkan daftar buku
func listBooks() {
	if len(books) == 0 {
		fmt.Println("No books available.")
		return
	}

	fmt.Println("List of books:")
	for _, book := range books {
		fmt.Printf("ID: %d, Title: %s, Author: %s, Year: %d\n", book.ID, book.Title, book.Author, book.Year)
	}
}
