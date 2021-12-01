package resources

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	validation "github.com/go-ozzo/ozzo-validation"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"HW7/internal/cache"
	"HW7/internal/models"
	"HW7/internal/store"
	"net/http"
)

type FactsResource struct {
	store store.FactsRepository
	cache cache.Cache
}

func NewFactsResource(store store.FactsRepository, cache cache.Cache) *FactsResource {
	return &FactsResource{
		store: store,
		cache: cache,
	}
}

func (cr *FactsResource) Routes() chi.Router {
	r := chi.NewRouter()

	r.Post("/", cr.CreateFact)
	r.Get("/", cr.AllFacts)
	r.Get("/{id}", cr.ById)
	r.Put("/", cr.UpdateFact)
	r.Delete("/{id}", cr.DeleteFact)

	return r
}

func (cr *FactsResource) CreateFact (w http.ResponseWriter, r *http.Request) {
	fact := new(models.Fact)
	fact.ID = primitive.NewObjectID()

	if err := json.NewDecoder(r.Body).Decode(fact); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Unknown err: %v", err)
		return
	}

	if err := cr.store.Create(r.Context(), fact); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "BD err: %v", err)
		return
	}

	if err := cr.cache.DeleteAll(r.Context()); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Cache err: %v", err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (cr *FactsResource) AllFacts(w http.ResponseWriter, r *http.Request) {
	queryValues := r.URL.Query()
	filter := &models.FactsFilter{}

	searchQuery := queryValues.Get("query")
	if searchQuery != "" {
		factsFromCache, err := cr.cache.Facts().Get(r.Context(), searchQuery)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Cache err: %v", err)
			return
		}
		if factsFromCache != nil {
			render.JSON(w, r, factsFromCache)
			return
		}

		filter.Query = &searchQuery
	}

	categories, err := cr.store.All(r.Context(), filter)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "DB err: %v", err)
		return
	}

	if searchQuery != "" && len(categories) > 0 {
		err = cr.cache.Facts().Set(r.Context(), searchQuery, categories)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Cache err: %v", err)
			return
		}
	}

	render.JSON(w, r, categories)
}

func (cr *FactsResource) ById(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	fact, err := cr.store.ByID(r.Context(), idStr)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "DB err: %v", err)
		return
	}

	render.JSON(w, r, fact)
}

func (cr *FactsResource) UpdateFact(w http.ResponseWriter, r *http.Request) {
	category := new(models.Fact)
	if err := json.NewDecoder(r.Body).Decode(category); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		fmt.Fprintf(w, "Unknown err: %v", err)
		return
	}

	err := validation.ValidateStruct(category,
		validation.Field(&category.ID, validation.Required),
		validation.Field(&category.Title, validation.Required))
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		fmt.Fprintf(w, "Unknown err: %v", err)
		return
	}

	if err = cr.store.Update(r.Context(), category); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "DB err: %v", err)
		return
	}
}

func (cr *FactsResource) DeleteFact(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	//id, err := strconv.Atoi(idStr)
	//if err != nil {
	//	w.WriteHeader(http.StatusInternalServerError)
	//	fmt.Fprintf(w, "Unknown err: %v", err)
	//	return
	//}

	if err := cr.store.Delete(r.Context(), idStr); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "DB err: %v", err)
		return
	}
}